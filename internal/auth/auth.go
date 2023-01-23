package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"mkuznets.com/go/sps/internal/user"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"net/http"
	"time"
)

type Service interface {
	Token(id string) (string, error)
	Middleware() func(next http.Handler) http.Handler
}

type authService struct {
	privateKey string
	publicKey  string

	privateKeyCache *rsa.PrivateKey
}

type jwtUser struct {
	id string
}

func (u *jwtUser) Id() string {
	return u.id
}

func New(privateKey, publicKey string) Service {
	return &authService{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

type claims struct {
	jwt.RegisteredClaims
	UserId string `json:"uid"`
}

func (s *authService) parsedPrivateKey() (*rsa.PrivateKey, error) {
	if s.privateKeyCache != nil {
		return s.privateKeyCache, nil
	}

	pemRaw, err := base64.StdEncoding.DecodeString(s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode private key: %w", err)
	}

	block, _ := pem.Decode(pemRaw)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM: `PRIVATE KEY` expected")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}
	if v, ok := key.(*rsa.PrivateKey); ok {
		s.privateKeyCache = v
		return v, nil
	}

	return nil, fmt.Errorf("RSA private key expected")
}

func (s *authService) Token(id string) (string, error) {
	privateKey, err := s.parsedPrivateKey()
	if err != nil {
		return "", err
	}

	cs := claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
		},
		UserId: id,
	}

	jwtEncoder := jwt.NewWithClaims(jwt.SigningMethodRS256, cs)

	token, err := jwtEncoder.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not create signed token: %w", err)
	}

	return token, nil
}

func (s *authService) keyFunc(_ *jwt.Token) (interface{}, error) {
	pemRaw, err := base64.StdEncoding.DecodeString(s.publicKey)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(pemRaw)
}

type cookieExtractor struct{}

func (e cookieExtractor) ExtractToken(req *http.Request) (string, error) {
	cookie, err := req.Cookie("JWT")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", request.ErrNoTokenInRequest
		}
		return "", err
	}

	return cookie.Value, nil
}

func (s *authService) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			extractor := request.MultiExtractor{}
			extractor = append(extractor, cookieExtractor{})
			extractor = append(extractor, request.BearerExtractor{})

			tokenString, err := extractor.ExtractToken(r)
			if err != nil {
				if err == request.ErrNoTokenInRequest {
					next.ServeHTTP(w, r)
					return
				}

				yerr.RenderJson(w, r, yerr.Unauthorised("valid auth header is required").WithError(err))
				return
			}

			var cs claims
			if _, err := jwt.ParseWithClaims(tokenString, &cs, s.keyFunc); err != nil {
				yerr.RenderJson(w, r, yerr.Unauthorised("invalid token").WithError(err))
				return
			}
			if !cs.VerifyExpiresAt(time.Now(), true) {
				yerr.RenderJson(w, r, yerr.Unauthorised("token expired"))
				return
			}

			ctx := user.Ctx(r.Context(), &jwtUser{id: cs.UserId})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
