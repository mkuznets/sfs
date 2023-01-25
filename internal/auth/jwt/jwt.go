package jwt

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/user"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
	"mkuznets.com/go/sfs/internal/ytils/yrender"
	"net/http"
	"time"
)

var (
	cookieName = "JWT"
)

type jwtService struct {
	privateKey string
	publicKey  string

	extractor       request.Extractor
	privateKeyCache *rsa.PrivateKey
}

func New(publicKey string) auth.Service {
	return &jwtService{
		privateKey: "",
		publicKey:  publicKey,
		extractor: request.MultiExtractor([]request.Extractor{
			request.BearerExtractor{},
			cookieExtractor{},
		}),
	}
}

func (s *jwtService) parsedPrivateKey() (*rsa.PrivateKey, error) {
	if s.privateKeyCache != nil {
		return s.privateKeyCache, nil
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.privateKey))
	if err != nil {
		return nil, yerr.New("could not parse private key").Err(err)
	}
	s.privateKeyCache = key
	return key, nil
}

func (s *jwtService) Token(id string) (string, error) {
	privateKey, err := s.parsedPrivateKey()
	if err != nil {
		return "", err
	}

	cs := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			Subject:   id,
		},
	}

	jwtEncoder := jwt.NewWithClaims(jwt.SigningMethodRS256, cs)

	token, err := jwtEncoder.SignedString(privateKey)
	if err != nil {
		return "", yerr.New("could not create signed token").Err(err)
	}

	return token, nil
}

func (s *jwtService) keyFunc(_ *jwt.Token) (interface{}, error) {
	return jwt.ParseRSAPublicKeyFromPEM([]byte(s.publicKey))
}

func (s *jwtService) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := s.extractor.ExtractToken(r)
			if err != nil {
				if err == request.ErrNoTokenInRequest {
					yrender.New(w, r, yerr.Unauthorised("JWT authentication required")).JSON()
					return
				}
				yrender.New(w, r, yerr.Unauthorised("valid JWT token is required").Err(err)).JSON()
				return
			}

			parser := jwt.NewParser(
				jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}),
			)

			var claims Claims
			if _, err := parser.ParseWithClaims(tokenString, &claims, s.keyFunc); err != nil {
				yrender.New(w, r, yerr.Unauthorised("invalid JWT token: %s", err)).JSON()
				return
			}

			ctx := user.Ctx(r.Context(), &claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
