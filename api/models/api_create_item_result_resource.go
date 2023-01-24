// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICreateItemResultResource api create item result resource
//
// swagger:model api.CreateItemResultResource
type APICreateItemResultResource struct {

	// id
	// Example: item_2K9BWVNuo3sG4yM322fbP3mB6ls
	ID string `json:"id,omitempty"`
}

// Validate validates this api create item result resource
func (m *APICreateItemResultResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api create item result resource based on context it is used
func (m *APICreateItemResultResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APICreateItemResultResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICreateItemResultResource) UnmarshalBinary(b []byte) error {
	var res APICreateItemResultResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}