package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"mkuznets.com/go/sfs/internal/user"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
	"mkuznets.com/go/sfs/internal/ytils/yrender"
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

func New(privateKey, publicKey string) Service {
	return &authService{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// Implements the `user.User` interface.
type customClaims struct {
	jwt.RegisteredClaims
}

func (c *customClaims) Id() string {
	return c.Subject
}

func (s *authService) parsedPrivateKey() (*rsa.PrivateKey, error) {
	if s.privateKeyCache != nil {
		return s.privateKeyCache, nil
	}

	pemRaw, err := base64.StdEncoding.DecodeString(s.privateKey)
	if err != nil {
		return nil, yerr.New("could not decode private key: %w", err)
	}

	block, _ := pem.Decode(pemRaw)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, yerr.New("failed to decode PEM: `PRIVATE KEY` expected")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, yerr.New("could not parse private key: %w", err)
	}
	if v, ok := key.(*rsa.PrivateKey); ok {
		s.privateKeyCache = v
		return v, nil
	}

	return nil, yerr.New("RSA private key expected")
}

func (s *authService) Token(id string) (string, error) {
	privateKey, err := s.parsedPrivateKey()
	if err != nil {
		return "", err
	}

	cs := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			Subject:   id,
		},
	}

	jwtEncoder := jwt.NewWithClaims(jwt.SigningMethodRS256, cs)

	token, err := jwtEncoder.SignedString(privateKey)
	if err != nil {
		return "", yerr.New("could not create signed token: %w", err)
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

				yrender.New(w, r, yerr.Unauthorised("valid auth header is required").Err(err)).JSON()
				return
			}

			var claims customClaims
			if _, err := jwt.ParseWithClaims(tokenString, &claims, s.keyFunc); err != nil {
				yrender.New(w, r, yerr.Unauthorised("invalid token").Err(err)).JSON()
				return
			}
			if !claims.VerifyExpiresAt(time.Now(), true) {
				yrender.New(w, r, yerr.Unauthorised("token expired")).JSON()
				return
			}

			ctx := user.Ctx(r.Context(), &claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
