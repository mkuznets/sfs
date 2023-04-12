package auth0

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"mkuznets.com/go/ytils/yhttp"

	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/user"
)

type auth0Service struct {
	issuerUrl *url.URL
	audience  string
}

func New(issuerUrl *url.URL, audience string) auth.Service {
	return &auth0Service{
		issuerUrl: issuerUrl,
		audience:  audience,
	}
}

func (s *auth0Service) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			provider := jwks.NewCachingProvider(s.issuerUrl, 5*time.Minute)
			jwtValidator, err := validator.New(
				provider.KeyFunc,
				validator.RS256,
				s.issuerUrl.String(),
				[]string{s.audience},
				validator.WithAllowedClockSkew(30*time.Second),
			)
			if err != nil {
				yhttp.Render(w, r, fmt.Errorf("HTTP 500: JWT validator error: %w", err)).JSON()
				return
			}

			token, err := jwtmiddleware.AuthHeaderTokenExtractor(r)
			if err != nil {
				yhttp.Render(w, r, fmt.Errorf("HTTP 401: auth token error: %w", err)).JSON()
				return
			}
			if token == "" {
				yhttp.Render(w, r, fmt.Errorf("HTTP 401: auth token is empty")).JSON()
				return
			}

			cs, err := jwtValidator.ValidateToken(r.Context(), token)
			if err != nil {
				yhttp.Render(w, r, fmt.Errorf("HTTP 401: auth token validation error: %w", err)).JSON()
				return
			}
			claims := cs.(*validator.ValidatedClaims)

			u := User{id: claims.RegisteredClaims.Subject}
			ctx := user.Ctx(r.Context(), &u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
