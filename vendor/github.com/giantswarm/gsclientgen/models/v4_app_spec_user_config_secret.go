// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4AppSpecUserConfigSecret v4 app spec user config secret
// swagger:model v4AppSpecUserConfigSecret
type V4AppSpecUserConfigSecret struct {

	// Name of the Secret on the control plane, which will become available wherever the app is installed
	Name string `json:"name,omitempty"`

	// Namespace of the Secret on the control plane, e.g. 123ab
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this v4 app spec user config secret
func (m *V4AppSpecUserConfigSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4AppSpecUserConfigSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AppSpecUserConfigSecret) UnmarshalBinary(b []byte) error {
	var res V4AppSpecUserConfigSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}