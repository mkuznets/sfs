// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the items client
type API interface {
	/*
	   CreateItems creates new items and returns a response with their i ds*/
	CreateItems(ctx context.Context, params *CreateItemsParams) (*CreateItemsOK, error)
	/*
	   GetItems gets items matching the given parameters*/
	GetItems(ctx context.Context, params *GetItemsParams) (*GetItemsOK, error)
}

// New creates a new items API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for items API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateItems creates new items and returns a response with their i ds
*/
func (a *Client) CreateItems(ctx context.Context, params *CreateItemsParams) (*CreateItemsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateItems",
		Method:             "POST",
		PathPattern:        "/items/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateItemsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateItemsOK), nil

}

/*
GetItems gets items matching the given parameters
*/
func (a *Client) GetItems(ctx context.Context, params *GetItemsParams) (*GetItemsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetItems",
		Method:             "POST",
		PathPattern:        "/items/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetItemsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetItemsOK), nil

}
