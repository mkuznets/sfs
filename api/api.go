package api

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"mkuznets.com/go/sps/api/client"
	"net/url"
)

type Api struct {
	client.Sps
	token string
}

// New create a new SPS API client instance.
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
		Sps:   *client.NewHTTPClientWithConfig(nil, &tc),
		token: token,
	}, nil
}

func (a *Api) AuthInfo() runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(a.token)
}

func NewAuth(key string) runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(key)
}
