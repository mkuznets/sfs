package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
	"strings"
	"time"
)

var TimeFunc = time.Now

// Implements the `user.User` interface.
type Claims struct {
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	now := TimeFunc()
	if !c.VerifyExpiresAt(now, true) {
		return yerr.New("token expired")
	}
	if !c.VerifyIssuedAt(now, true) {
		return yerr.New("token used before issued")
	}
	if strings.TrimSpace(c.Subject) == "" {
		return yerr.New("token subject is empty")
	}

	return nil
}

func (c *Claims) Id() string {
	return c.Subject
}
