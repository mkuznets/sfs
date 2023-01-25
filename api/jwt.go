package api

import (
	"crypto/rsa"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

var m sync.Mutex

// Implements runtime.ClientAuthInfoWriter
type jwtAuthenticator struct {
	subject    string
	pk         string
	pkCache    *rsa.PrivateKey
	tokenCache string
}

func (j *jwtAuthenticator) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	token, err := j.token()
	if err != nil {
		return err
	}
	return req.SetHeaderParam("Authorization", "Bearer "+token)
}

func newJwtAuthenticator(privateKey, subject string) runtime.ClientAuthInfoWriter {
	return &jwtAuthenticator{
		subject: subject,
		pk:      privateKey,
	}
}

func (j *jwtAuthenticator) parsePrivateKey() (*rsa.PrivateKey, error) {
	m.Lock()
	defer m.Unlock()

	if j.pkCache != nil {
		return j.pkCache, nil
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(j.pk))
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %s", err)
	}

	j.pkCache = key
	return key, nil
}

func (j *jwtAuthenticator) token() (string, error) {
	privateKey, err := j.parsePrivateKey()
	if err != nil {
		return "", err
	}

	m.Lock()
	defer m.Unlock()

	if j.tokenCache != "" {
		return j.tokenCache, nil
	}

	cs := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
		Subject:   j.subject,
	}

	jwtEncoder := jwt.NewWithClaims(jwt.SigningMethodRS256, cs)

	token, err := jwtEncoder.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not create signed token: %w", err)
	}

	j.tokenCache = token

	return token, nil
}
