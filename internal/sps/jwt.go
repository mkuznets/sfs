package sps

import (
	"encoding/base64"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"net/http"
	"time"
)

var expirationLimit = 30 * time.Minute

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
				renderApiError(w, r, err, http.StatusUnauthorized, "valid auth header is required")
				return
			}

			var claims jwt.RegisteredClaims
			if _, err := jwt.ParseWithClaims(tokenString, &claims, j.keyFunc); err != nil {
				renderApiError(w, r, err, http.StatusUnauthorized, "invalid token: "+err.Error())
				return
			}
			if !claims.VerifyExpiresAt(time.Now().Add(-expirationLimit), true) {
				renderApiError(w, r, nil, http.StatusUnauthorized, "token expired")
				return
			}
			if !claims.VerifyIssuedAt(time.Now().Add(-expirationLimit), true) {
				renderApiError(w, r, nil, http.StatusUnauthorized, "token expired")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
