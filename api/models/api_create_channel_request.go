// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICreateChannelRequest api create channel request
//
// swagger:model api.CreateChannelRequest
type APICreateChannelRequest struct {

	// authors
	// Example: The Owl
	Authors string `json:"authors,omitempty"`

	// description
	// Example: Bored owls talk about whatever happens to be on their minds
	Description string `json:"description,omitempty"`

	// link
	// Example: https://example.com
	Link string `json:"link,omitempty"`

	// title
	// Example: Bored Owls Online Radio
	Title string `json:"title,omitempty"`
}

// Validate validates this api create channel request
func (m *APICreateChannelRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api create channel request based on context it is used
func (m *APICreateChannelRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APICreateChannelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICreateChannelRequest) UnmarshalBinary(b []byte) error {
	var res APICreateChannelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
