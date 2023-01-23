// Code generated by go-swagger; DO NOT EDIT.

package episodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new episodes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for episodes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateItems(params *CreateItemsParams, opts ...ClientOption) (*CreateItemsOK, error)

	GetItems(params *GetItemsParams, opts ...ClientOption) (*GetItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateItems creates new episodes and returns a response with their i ds
*/
func (a *Client) CreateItems(params *CreateItemsParams, opts ...ClientOption) (*CreateItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateItems",
		Method:             "POST",
		PathPattern:        "/episodes/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateItemsReader{formats: a.formats},
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
	success, ok := result.(*CreateItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetItems gets episodes matching the given parameters
*/
func (a *Client) GetItems(params *GetItemsParams, opts ...ClientOption) (*GetItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetItems",
		Method:             "POST",
		PathPattern:        "/episodes/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetItemsReader{formats: a.formats},
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
	success, ok := result.(*GetItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
