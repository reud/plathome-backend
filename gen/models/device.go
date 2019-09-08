// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Device device
// swagger:model Device
type Device struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// ez requester models
	// Required: true
	EzRequesterModels []*EzRequesterModel `json:"ezRequesterModels"`

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEzRequesterModels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateEzRequesterModels(formats strfmt.Registry) error {

	if err := validate.Required("ezRequesterModels", "body", m.EzRequesterModels); err != nil {
		return err
	}

	for i := 0; i < len(m.EzRequesterModels); i++ {
		if swag.IsZero(m.EzRequesterModels[i]) { // not required
			continue
		}

		if m.EzRequesterModels[i] != nil {
			if err := m.EzRequesterModels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ezRequesterModels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
