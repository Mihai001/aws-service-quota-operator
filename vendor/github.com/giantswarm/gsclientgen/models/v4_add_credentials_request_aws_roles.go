// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V4AddCredentialsRequestAwsRoles IAM roles to assume by certain entities
// swagger:model v4AddCredentialsRequestAwsRoles
type V4AddCredentialsRequestAwsRoles struct {

	// ARN of the IAM role to assume by Giant Swarm support staff
	// Required: true
	Admin *string `json:"admin"`

	// ARN of the IAM role to assume by the software operating clusters
	// Required: true
	Awsoperator *string `json:"awsoperator"`
}

// Validate validates this v4 add credentials request aws roles
func (m *V4AddCredentialsRequestAwsRoles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsoperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4AddCredentialsRequestAwsRoles) validateAdmin(formats strfmt.Registry) error {

	if err := validate.Required("admin", "body", m.Admin); err != nil {
		return err
	}

	return nil
}

func (m *V4AddCredentialsRequestAwsRoles) validateAwsoperator(formats strfmt.Registry) error {

	if err := validate.Required("awsoperator", "body", m.Awsoperator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4AddCredentialsRequestAwsRoles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4AddCredentialsRequestAwsRoles) UnmarshalBinary(b []byte) error {
	var res V4AddCredentialsRequestAwsRoles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}