// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemFileResource item file resource
//
// swagger:model ItemFileResource
type ItemFileResource struct {
	// content type
	// Example: audio/mpeg
	ContentType string `json:"content_type,omitempty"`

	// id
	// Example: file_2K9BWVNuo3sG4yM322fbP3mB6ls
	ID string `json:"id,omitempty"`

	// size
	// Example: 123456
	Size int64 `json:"size,omitempty"`

	// url
	// Example: https://example.com/file.mp3
	URL string `json:"url,omitempty"`
}

// Validate validates this item file resource
func (m *ItemFileResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this item file resource based on context it is used
func (m *ItemFileResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemFileResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemFileResource) UnmarshalBinary(b []byte) error {
	var res ItemFileResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
