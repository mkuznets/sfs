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

// ItemResource item resource
//
// swagger:model ItemResource
type ItemResource struct {
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

	// feed id
	// Example: feed_2K9BWVNuo3sG4yM322fbP3mB6ls
	FeedID string `json:"feed_id,omitempty"`

	// file
	File struct {
		ItemFileResource
	} `json:"file,omitempty"`

	// id
	// Example: item_2K9BWVNuo3sG4yM322fbP3mB6ls
	ID string `json:"id,omitempty"`

	// link
	// Example: https://example.com
	Link string `json:"link,omitempty"`

	// published at
	// Example: 2023-01-01T01:02:03.456Z
	// Format: date-time
	PublishedAt strfmt.DateTime `json:"published_at,omitempty"`

	// title
	// Example: Bored Owls Online Radio
	Title string `json:"title,omitempty"`

	// updated at
	// Example: 2023-01-01T01:02:03.456Z
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this item resource
func (m *ItemResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
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

func (m *ItemResource) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ItemResource) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.File) { // not required
		return nil
	}

	return nil
}

func (m *ItemResource) validatePublishedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("published_at", "body", "date-time", m.PublishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ItemResource) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this item resource based on the context it is used
func (m *ItemResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemResource) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemResource) UnmarshalBinary(b []byte) error {
	var res ItemResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
