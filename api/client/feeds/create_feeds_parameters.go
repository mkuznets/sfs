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

// NewCreateFeedsParams creates a new CreateFeedsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFeedsParams() *CreateFeedsParams {
	return &CreateFeedsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFeedsParamsWithTimeout creates a new CreateFeedsParams object
// with the ability to set a timeout on a request.
func NewCreateFeedsParamsWithTimeout(timeout time.Duration) *CreateFeedsParams {
	return &CreateFeedsParams{
		timeout: timeout,
	}
}

// NewCreateFeedsParamsWithContext creates a new CreateFeedsParams object
// with the ability to set a context for a request.
func NewCreateFeedsParamsWithContext(ctx context.Context) *CreateFeedsParams {
	return &CreateFeedsParams{
		Context: ctx,
	}
}

// NewCreateFeedsParamsWithHTTPClient creates a new CreateFeedsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFeedsParamsWithHTTPClient(client *http.Client) *CreateFeedsParams {
	return &CreateFeedsParams{
		HTTPClient: client,
	}
}

/*
CreateFeedsParams contains all the parameters to send to the API endpoint

	for the create feeds operation.

	Typically these are written to a http.Request.
*/
type CreateFeedsParams struct {

	/* Request.

	   CreateFeeds request
	*/
	Request *models.CreateFeedsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create feeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFeedsParams) WithDefaults() *CreateFeedsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create feeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFeedsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create feeds params
func (o *CreateFeedsParams) WithTimeout(timeout time.Duration) *CreateFeedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create feeds params
func (o *CreateFeedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create feeds params
func (o *CreateFeedsParams) WithContext(ctx context.Context) *CreateFeedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create feeds params
func (o *CreateFeedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create feeds params
func (o *CreateFeedsParams) WithHTTPClient(client *http.Client) *CreateFeedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create feeds params
func (o *CreateFeedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create feeds params
func (o *CreateFeedsParams) WithRequest(request *models.CreateFeedsRequest) *CreateFeedsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create feeds params
func (o *CreateFeedsParams) SetRequest(request *models.CreateFeedsRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFeedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
