// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"mkuznets.com/go/sfs/api/client/feeds"
	"mkuznets.com/go/sfs/api/client/files"
	"mkuznets.com/go/sfs/api/client/items"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// New creates a new simple feed service HTTP client.
func New(c Config) *SimpleFeedService {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(SimpleFeedService)
	cli.Transport = transport
	cli.Feeds = feeds.New(transport, strfmt.Default, c.AuthInfo)
	cli.Files = files.New(transport, strfmt.Default, c.AuthInfo)
	cli.Items = items.New(transport, strfmt.Default, c.AuthInfo)
	return cli
}

// SimpleFeedService is a client for simple feed service
type SimpleFeedService struct {
	Feeds     *feeds.Client
	Files     *files.Client
	Items     *items.Client
	Transport runtime.ClientTransport
}
