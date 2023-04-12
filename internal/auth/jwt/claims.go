package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var TimeFunc = time.Now

// Implements the `user.User` interface.
type Claims struct {
	jwt.RegisteredClaims
}

func (c *Claims) Valid() error {
	now := TimeFunc()
	if !c.VerifyExpiresAt(now, true) {
		return errors.New("token expired")
	}
	if !c.VerifyIssuedAt(now, true) {
		return errors.New("token used before issued")
	}
	if strings.TrimSpace(c.Subject) == "" {
		return errors.New("token subject is empty")
	}

	return nil
}

func (c *Claims) Id() string {
	return c.Subject
}
