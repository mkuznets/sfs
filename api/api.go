package api

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"mkuznets.com/go/sfs/api/client"
	"net/url"
)

type Api struct {
	client.SimpleFeedService
	token string
}

// New create a new Simple Feed Service API client instance.
func New(baseUrl, token string) (*Api, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	tc := client.TransportConfig{
		Host:     u.Host,
		BasePath: u.Path,
		Schemes:  []string{u.Scheme},
	}

	return &Api{
		SimpleFeedService: *client.NewHTTPClientWithConfig(nil, &tc),
		token:             token,
	}, nil
}

func (a *Api) AuthInfo() runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(a.token)
}

func NewAuth(key string) runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(key)
}
