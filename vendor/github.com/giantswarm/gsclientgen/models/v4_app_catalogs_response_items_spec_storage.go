// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4AppCatalogsResponseItemsSpecStorage Contains information on where to find the full catalog, and what type of catalog it is.
// swagger:model v4AppCatalogsResponseItemsSpecStorage
type V4AppCatalogsResponseItemsSpecStorage struct {

	// A URL where clients can download the full catalog.
	URL string `json:"URL,omitempty"`

	// The format of this catalog. (Currently only helm is supported.)
	Type string `json:"type,omitempty"`
}

// Validate validates this v4 app catalogs response items spec storage
func (m *V4AppCatalogsResponseItemsSpecStorage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4AppCatalogsResponseItemsSpecStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AppCatalogsResponseItemsSpecStorage) UnmarshalBinary(b []byte) error {
	var res V4AppCatalogsResponseItemsSpecStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}