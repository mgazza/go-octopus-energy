// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountGasMeterPoint account gas meter point
//
// swagger:model AccountGasMeterPoint
type AccountGasMeterPoint struct {

	// agreements
	Agreements []*AccountAgreement `json:"agreements"`

	// consumption standard
	// Example: 3448
	ConsumptionStandard int64 `json:"consumption_standard,omitempty"`

	// meters
	Meters []*GasMeter `json:"meters"`

	// mprn
	// Example: 1234567890
	Mprn string `json:"mprn,omitempty"`
}

// Validate validates this account gas meter point
func (m *AccountGasMeterPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgreements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountGasMeterPoint) validateAgreements(formats strfmt.Registry) error {
	if swag.IsZero(m.Agreements) { // not required
		return nil
	}

	for i := 0; i < len(m.Agreements); i++ {
		if swag.IsZero(m.Agreements[i]) { // not required
			continue
		}

		if m.Agreements[i] != nil {
			if err := m.Agreements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agreements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agreements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountGasMeterPoint) validateMeters(formats strfmt.Registry) error {
	if swag.IsZero(m.Meters) { // not required
		return nil
	}

	for i := 0; i < len(m.Meters); i++ {
		if swag.IsZero(m.Meters[i]) { // not required
			continue
		}

		if m.Meters[i] != nil {
			if err := m.Meters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("meters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("meters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account gas meter point based on the context it is used
func (m *AccountGasMeterPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgreements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountGasMeterPoint) contextValidateAgreements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Agreements); i++ {

		if m.Agreements[i] != nil {

			if swag.IsZero(m.Agreements[i]) { // not required
				return nil
			}

			if err := m.Agreements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agreements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agreements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountGasMeterPoint) contextValidateMeters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Meters); i++ {

		if m.Meters[i] != nil {

			if swag.IsZero(m.Meters[i]) { // not required
				return nil
			}

			if err := m.Meters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("meters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("meters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountGasMeterPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountGasMeterPoint) UnmarshalBinary(b []byte) error {
	var res AccountGasMeterPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
