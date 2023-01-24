// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UploadFileResultResource upload file result resource
//
// swagger:model UploadFileResultResource
type UploadFileResultResource struct {

	// error
	// Example: invalid file format
	Error string `json:"error,omitempty"`

	// id
	// Example: file_2K9BWVNuo3sG4yM322fbP3mB6ls
	ID string `json:"id,omitempty"`
}

// Validate validates this upload file result resource
func (m *UploadFileResultResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upload file result resource based on context it is used
func (m *UploadFileResultResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadFileResultResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadFileResultResource) UnmarshalBinary(b []byte) error {
	var res UploadFileResultResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}