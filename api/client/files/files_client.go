// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the files client
type API interface {
	/*
	   UploadFiles uploads new audio files*/
	UploadFiles(ctx context.Context, params *UploadFilesParams) (*UploadFilesOK, error)
}

// New creates a new files API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for files API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
UploadFiles uploads new audio files
*/
func (a *Client) UploadFiles(ctx context.Context, params *UploadFilesParams) (*UploadFilesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadFiles",
		Method:             "POST",
		PathPattern:        "/files/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadFilesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadFilesOK), nil

}
