// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FeedResource feed resource
//
// swagger:model FeedResource
type FeedResource struct {
	// authors
	// Example: The Owl
	Authors string `json:"authors,omitempty"`

	// created at
	// Example: 2023-01-01T01:02:03.456Z
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	// Example: Bored owls talk about whatever happens to be on their minds
	Description string `json:"description,omitempty"`

	// id
	// Example: feed_2K9BWVNuo3sG4yM322fbP3mB6ls
	ID string `json:"id,omitempty"`

	// link
	// Example: https://example.com
	Link string `json:"link,omitempty"`

	// rss url
	// Example: https://example.com/feed.rss
	RssURL string `json:"rss_url,omitempty"`

	// title
	// Example: Bored Owls Online Radio
	Title string `json:"title,omitempty"`

	// updated at
	// Example: 2023-01-01T01:02:03.456Z
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this feed resource
func (m *FeedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeedResource) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FeedResource) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this feed resource based on context it is used
func (m *FeedResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedResource) UnmarshalBinary(b []byte) error {
	var res FeedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
