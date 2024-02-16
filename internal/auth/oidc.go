package auth

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

var (
	ErrInternal   = errors.New("internal error")
	ErrValidation = errors.New("could not validate token")
)

type oidcService struct {
	issuerURL    *url.URL
	audience     string
	jwksProvider *jwks.CachingProvider
}

type oidcUser struct {
	id string
}

func (c *oidcUser) ID() string {
	return c.id
}

func NewOIDCService(issuerURL *url.URL, audience string) Service {
	return &oidcService{
		issuerURL:    issuerURL,
		audience:     audience,
		jwksProvider: jwks.NewCachingProvider(issuerURL, 5*time.Minute),
	}
}

func (s *oidcService) Middleware(handleError func(w http.ResponseWriter, r *http.Request, err error)) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtValidator, err := validator.New(
				s.jwksProvider.KeyFunc,
				validator.RS256,
				s.issuerURL.String(),
				[]string{s.audience},
				validator.WithAllowedClockSkew(30*time.Second),
			)
			if err != nil {
				handleError(w, r, fmt.Errorf("%w: %w", ErrInternal, err))
				return
			}

			token, err := jwtmiddleware.AuthHeaderTokenExtractor(r)
			if err != nil {
				handleError(w, r, fmt.Errorf("%w: %w", ErrValidation, err))
				return
			}
			if token == "" {
				handleError(w, r, fmt.Errorf("%w: token is empty", ErrValidation))
				return
			}

			cs, err := jwtValidator.ValidateToken(r.Context(), token)
			if err != nil {
				handleError(w, r, fmt.Errorf("%w: %w", ErrValidation, err))
				return
			}
			claims := cs.(*validator.ValidatedClaims)

			u := oidcUser{id: claims.RegisteredClaims.Subject}
			ctx := NewUserContext(r.Context(), &u)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
