package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

const (
	expiresIn = 2 * 365 * 24 * time.Hour
)

var m sync.Mutex

// Implements runtime.ClientAuthInfoWriter
type jwtAuthenticator struct {
	userId     string
	pk         *rsa.PrivateKey
	tokenCache string
}

func (j *jwtAuthenticator) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	token, err := j.token()
	if err != nil {
		return err
	}
	return req.SetHeaderParam("Authorization", "Bearer "+token)
}

func NewJwtAuthInfo(privateKey *rsa.PrivateKey, userId string) runtime.ClientAuthInfoWriter {
	return &jwtAuthenticator{
		userId: userId,
		pk:     privateKey,
	}
}

func (j *jwtAuthenticator) token() (string, error) {
	if j.tokenCache != "" {
		return j.tokenCache, nil
	}

	m.Lock()
	defer m.Unlock()

	cs := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject:   j.userId,
	}

	jwtEncoder := jwt.NewWithClaims(jwt.SigningMethodRS256, cs)

	token, err := jwtEncoder.SignedString(j.pk)
	if err != nil {
		return "", fmt.Errorf("could not create signed token: %w", err)
	}

	j.tokenCache = token

	return token, nil
}
