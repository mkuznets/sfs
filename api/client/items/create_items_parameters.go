// Code generated by go-swagger; DO NOT EDIT.

package items

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

// NewCreateItemsParams creates a new CreateItemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateItemsParams() *CreateItemsParams {
	return &CreateItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateItemsParamsWithTimeout creates a new CreateItemsParams object
// with the ability to set a timeout on a request.
func NewCreateItemsParamsWithTimeout(timeout time.Duration) *CreateItemsParams {
	return &CreateItemsParams{
		timeout: timeout,
	}
}

// NewCreateItemsParamsWithContext creates a new CreateItemsParams object
// with the ability to set a context for a request.
func NewCreateItemsParamsWithContext(ctx context.Context) *CreateItemsParams {
	return &CreateItemsParams{
		Context: ctx,
	}
}

// NewCreateItemsParamsWithHTTPClient creates a new CreateItemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateItemsParamsWithHTTPClient(client *http.Client) *CreateItemsParams {
	return &CreateItemsParams{
		HTTPClient: client,
	}
}

/*
CreateItemsParams contains all the parameters to send to the API endpoint

	for the create items operation.

	Typically these are written to a http.Request.
*/
type CreateItemsParams struct {
	/* Request.

	   CreateItems request
	*/
	Request *models.CreateItemsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateItemsParams) WithDefaults() *CreateItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create items params
func (o *CreateItemsParams) WithTimeout(timeout time.Duration) *CreateItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create items params
func (o *CreateItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create items params
func (o *CreateItemsParams) WithContext(ctx context.Context) *CreateItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create items params
func (o *CreateItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create items params
func (o *CreateItemsParams) WithHTTPClient(client *http.Client) *CreateItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create items params
func (o *CreateItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create items params
func (o *CreateItemsParams) WithRequest(request *models.CreateItemsRequest) *CreateItemsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create items params
func (o *CreateItemsParams) SetRequest(request *models.CreateItemsRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
