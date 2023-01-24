// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetChannelsIDParams creates a new GetChannelsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChannelsIDParams() *GetChannelsIDParams {
	return &GetChannelsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChannelsIDParamsWithTimeout creates a new GetChannelsIDParams object
// with the ability to set a timeout on a request.
func NewGetChannelsIDParamsWithTimeout(timeout time.Duration) *GetChannelsIDParams {
	return &GetChannelsIDParams{
		timeout: timeout,
	}
}

// NewGetChannelsIDParamsWithContext creates a new GetChannelsIDParams object
// with the ability to set a context for a request.
func NewGetChannelsIDParamsWithContext(ctx context.Context) *GetChannelsIDParams {
	return &GetChannelsIDParams{
		Context: ctx,
	}
}

// NewGetChannelsIDParamsWithHTTPClient creates a new GetChannelsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChannelsIDParamsWithHTTPClient(client *http.Client) *GetChannelsIDParams {
	return &GetChannelsIDParams{
		HTTPClient: client,
	}
}

/*
GetChannelsIDParams contains all the parameters to send to the API endpoint

	for the get channels ID operation.

	Typically these are written to a http.Request.
*/
type GetChannelsIDParams struct {

	/* ID.

	   Feed ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get channels ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChannelsIDParams) WithDefaults() *GetChannelsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get channels ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChannelsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get channels ID params
func (o *GetChannelsIDParams) WithTimeout(timeout time.Duration) *GetChannelsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get channels ID params
func (o *GetChannelsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get channels ID params
func (o *GetChannelsIDParams) WithContext(ctx context.Context) *GetChannelsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get channels ID params
func (o *GetChannelsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get channels ID params
func (o *GetChannelsIDParams) WithHTTPClient(client *http.Client) *GetChannelsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get channels ID params
func (o *GetChannelsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get channels ID params
func (o *GetChannelsIDParams) WithID(id string) *GetChannelsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get channels ID params
func (o *GetChannelsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetChannelsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}