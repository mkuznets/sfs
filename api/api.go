package api

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"mkuznets.com/go/sfs/api/client"
)

type Api struct {
	client.SimpleFeedService
	token string
	auth  runtime.ClientAuthInfoWriter
}

// New create a new Simple Feed Service API client instance.
func New(scheme, host, basePath string) *Api {
	tc := client.TransportConfig{
		Host:     host,
		BasePath: basePath,
		Schemes:  []string{scheme},
	}

	return &Api{
		SimpleFeedService: *client.NewHTTPClientWithConfig(nil, &tc),
		auth:              httptransport.PassThroughAuth,
	}
}

func (a *Api) WithJwt(privateKey, userId string) *Api {
	a.auth = newJwtAuthenticator(privateKey, userId)
	return a
}

func (a *Api) AuthWriter() runtime.ClientAuthInfoWriter {
	return a.auth
}
