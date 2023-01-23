// Code generated by go-swagger; DO NOT EDIT.

package episodes_page

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new episodes page API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for episodes page API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEpisode(params *CreateEpisodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEpisodeOK, error)

	ListEpisodes(params *ListEpisodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListEpisodesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateEpisode creates a new episode
*/
func (a *Client) CreateEpisode(params *CreateEpisodeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateEpisodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEpisodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateEpisode",
		Method:             "POST",
		PathPattern:        "/channels/{id}/episodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateEpisodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateEpisodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateEpisode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListEpisodes lists episoded of the given channel
*/
func (a *Client) ListEpisodes(params *ListEpisodesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListEpisodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEpisodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListEpisodes",
		Method:             "GET",
		PathPattern:        "/channels/{id}/episodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListEpisodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListEpisodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListEpisodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
