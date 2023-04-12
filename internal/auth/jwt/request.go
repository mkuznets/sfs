package jwt

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4/request"
)

type cookieExtractor struct{}

func (e cookieExtractor) ExtractToken(req *http.Request) (string, error) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", request.ErrNoTokenInRequest
		}
		return "", err
	}

	return cookie.Value, nil
}
