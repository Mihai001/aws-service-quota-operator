// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4InfoResponseFeaturesNodepools Support for grouping of worker nodes into node pools.
// swagger:model v4InfoResponseFeaturesNodepools
type V4InfoResponseFeaturesNodepools struct {

	// The minimum release version number required to have support for node pools.
	ReleaseVersionMinimum string `json:"release_version_minimum,omitempty"`
}

// Validate validates this v4 info response features nodepools
func (m *V4InfoResponseFeaturesNodepools) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4InfoResponseFeaturesNodepools) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4InfoResponseFeaturesNodepools) UnmarshalBinary(b []byte) error {
	var res V4InfoResponseFeaturesNodepools
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}