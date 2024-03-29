// Code generated by go-swagger; DO NOT EDIT.

package feeds

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

	"mkuznets.com/go/sfs/api/models"
)

// NewGetFeedsParams creates a new GetFeedsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFeedsParams() *GetFeedsParams {
	return &GetFeedsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeedsParamsWithTimeout creates a new GetFeedsParams object
// with the ability to set a timeout on a request.
func NewGetFeedsParamsWithTimeout(timeout time.Duration) *GetFeedsParams {
	return &GetFeedsParams{
		timeout: timeout,
	}
}

// NewGetFeedsParamsWithContext creates a new GetFeedsParams object
// with the ability to set a context for a request.
func NewGetFeedsParamsWithContext(ctx context.Context) *GetFeedsParams {
	return &GetFeedsParams{
		Context: ctx,
	}
}

// NewGetFeedsParamsWithHTTPClient creates a new GetFeedsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFeedsParamsWithHTTPClient(client *http.Client) *GetFeedsParams {
	return &GetFeedsParams{
		HTTPClient: client,
	}
}

/*
GetFeedsParams contains all the parameters to send to the API endpoint

	for the get feeds operation.

	Typically these are written to a http.Request.
*/
type GetFeedsParams struct {
	/* Request.

	   Parameters for filtering feeds
	*/
	Request *models.GetFeedsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get feeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeedsParams) WithDefaults() *GetFeedsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get feeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeedsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get feeds params
func (o *GetFeedsParams) WithTimeout(timeout time.Duration) *GetFeedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feeds params
func (o *GetFeedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feeds params
func (o *GetFeedsParams) WithContext(ctx context.Context) *GetFeedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feeds params
func (o *GetFeedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feeds params
func (o *GetFeedsParams) WithHTTPClient(client *http.Client) *GetFeedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feeds params
func (o *GetFeedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the get feeds params
func (o *GetFeedsParams) WithRequest(request *models.GetFeedsRequest) *GetFeedsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the get feeds params
func (o *GetFeedsParams) SetRequest(request *models.GetFeedsRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
