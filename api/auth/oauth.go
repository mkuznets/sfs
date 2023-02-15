package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"golang.org/x/oauth2"
)

// Implements runtime.ClientAuthInfoWriter
type oauthAuthenticator struct {
	tokenSource oauth2.TokenSource
}

func (j *oauthAuthenticator) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	t, err := j.tokenSource.Token()
	if err != nil {
		return err
	}
	return req.SetHeaderParam("Authorization", fmt.Sprintf("%s %s", t.Type(), t.AccessToken))
}

func NewOauthInfo(privateKey *rsa.PrivateKey, userId string) runtime.ClientAuthInfoWriter {
	return &jwtAuthenticator{
		userId: userId,
		pk:     privateKey,
	}
}
