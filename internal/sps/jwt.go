package sps

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"net/http"
	"time"
)

var (
	expirationLimit = 30 * time.Minute
)

type Jwt struct {
	PublicKey string `long:"public-key" description:"RSA public key" required:"true"`
}

func (j *Jwt) keyFunc(token *jwt.Token) (interface{}, error) {
	pemRaw, err := base64.StdEncoding.DecodeString(j.PublicKey)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(pemRaw)
}

func (j *Jwt) Validator() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := request.BearerExtractor{}.ExtractToken(r)
			if err != nil {
				yerr.RenderJson(w, r, yerr.Unauthorised("valid auth header is required").WithError(err))
				return
			}

			var claims jwt.RegisteredClaims
			if _, err := jwt.ParseWithClaims(tokenString, &claims, j.keyFunc); err != nil {
				yerr.RenderJson(w, r, yerr.Unauthorised("invalid token").WithError(err))
				return
			}
			if !claims.VerifyExpiresAt(time.Now().Add(-expirationLimit), true) {
				yerr.RenderJson(w, r, yerr.Unauthorised("token expired"))
				return
			}
			if !claims.VerifyIssuedAt(time.Now().Add(-expirationLimit), true) {
				yerr.RenderJson(w, r, yerr.Unauthorised("token expired"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
