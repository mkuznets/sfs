// Code generated by go-swagger; DO NOT EDIT.

package episodes_page

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

	"mkuznets.com/go/sps/api/models"
)

// NewCreateEpisodesParams creates a new CreateEpisodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEpisodesParams() *CreateEpisodesParams {
	return &CreateEpisodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEpisodesParamsWithTimeout creates a new CreateEpisodesParams object
// with the ability to set a timeout on a request.
func NewCreateEpisodesParamsWithTimeout(timeout time.Duration) *CreateEpisodesParams {
	return &CreateEpisodesParams{
		timeout: timeout,
	}
}

// NewCreateEpisodesParamsWithContext creates a new CreateEpisodesParams object
// with the ability to set a context for a request.
func NewCreateEpisodesParamsWithContext(ctx context.Context) *CreateEpisodesParams {
	return &CreateEpisodesParams{
		Context: ctx,
	}
}

// NewCreateEpisodesParamsWithHTTPClient creates a new CreateEpisodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEpisodesParamsWithHTTPClient(client *http.Client) *CreateEpisodesParams {
	return &CreateEpisodesParams{
		HTTPClient: client,
	}
}

/*
CreateEpisodesParams contains all the parameters to send to the API endpoint

	for the create episodes operation.

	Typically these are written to a http.Request.
*/
type CreateEpisodesParams struct {

	/* ID.

	   Feed ID
	*/
	ID string

	/* Request.

	   CreateItems request
	*/
	Request *models.CreateEpisodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create episodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEpisodesParams) WithDefaults() *CreateEpisodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create episodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEpisodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create episodes params
func (o *CreateEpisodesParams) WithTimeout(timeout time.Duration) *CreateEpisodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create episodes params
func (o *CreateEpisodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create episodes params
func (o *CreateEpisodesParams) WithContext(ctx context.Context) *CreateEpisodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create episodes params
func (o *CreateEpisodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create episodes params
func (o *CreateEpisodesParams) WithHTTPClient(client *http.Client) *CreateEpisodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create episodes params
func (o *CreateEpisodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create episodes params
func (o *CreateEpisodesParams) WithID(id string) *CreateEpisodesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create episodes params
func (o *CreateEpisodesParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the create episodes params
func (o *CreateEpisodesParams) WithRequest(request *models.CreateEpisodeRequest) *CreateEpisodesParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create episodes params
func (o *CreateEpisodesParams) SetRequest(request *models.CreateEpisodeRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEpisodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
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