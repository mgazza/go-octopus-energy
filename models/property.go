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
	"github.com/go-openapi/validate"
)

// Property property
//
// swagger:model Property
type Property struct {

	// address line 1
	// Example: 10 Downing Street
	AddressLine1 string `json:"address_line_1,omitempty"`

	// address line 2
	AddressLine2 string `json:"address_line_2,omitempty"`

	// address line 3
	AddressLine3 string `json:"address_line_3,omitempty"`

	// county
	County string `json:"county,omitempty"`

	// electricity meter points
	ElectricityMeterPoints []*AccountElectricityMeterPoint `json:"electricity_meter_points"`

	// gas meter points
	GasMeterPoints []*AccountGasMeterPoint `json:"gas_meter_points"`

	// id
	// Example: 1234567
	ID int64 `json:"id,omitempty"`

	// moved in at
	// Example: 2020-11-30T00:00:00Z
	// Format: datetime
	MovedInAt strfmt.DateTime `json:"moved_in_at,omitempty"`

	// moved out at
	// Format: datetime
	MovedOutAt strfmt.DateTime `json:"moved_out_at,omitempty"`

	// postcode
	// Example: W1 1AA
	Postcode string `json:"postcode,omitempty"`

	// town
	// Example: LONDON
	Town string `json:"town,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElectricityMeterPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasMeterPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMovedInAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMovedOutAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) validateElectricityMeterPoints(formats strfmt.Registry) error {
	if swag.IsZero(m.ElectricityMeterPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.ElectricityMeterPoints); i++ {
		if swag.IsZero(m.ElectricityMeterPoints[i]) { // not required
			continue
		}

		if m.ElectricityMeterPoints[i] != nil {
			if err := m.ElectricityMeterPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("electricity_meter_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("electricity_meter_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Property) validateGasMeterPoints(formats strfmt.Registry) error {
	if swag.IsZero(m.GasMeterPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.GasMeterPoints); i++ {
		if swag.IsZero(m.GasMeterPoints[i]) { // not required
			continue
		}

		if m.GasMeterPoints[i] != nil {
			if err := m.GasMeterPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gas_meter_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gas_meter_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Property) validateMovedInAt(formats strfmt.Registry) error {
	if swag.IsZero(m.MovedInAt) { // not required
		return nil
	}

	if err := validate.FormatOf("moved_in_at", "body", "datetime", m.MovedInAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Property) validateMovedOutAt(formats strfmt.Registry) error {
	if swag.IsZero(m.MovedOutAt) { // not required
		return nil
	}

	if err := validate.FormatOf("moved_out_at", "body", "datetime", m.MovedOutAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this property based on the context it is used
func (m *Property) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElectricityMeterPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGasMeterPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) contextValidateElectricityMeterPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ElectricityMeterPoints); i++ {

		if m.ElectricityMeterPoints[i] != nil {

			if swag.IsZero(m.ElectricityMeterPoints[i]) { // not required
				return nil
			}

			if err := m.ElectricityMeterPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("electricity_meter_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("electricity_meter_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Property) contextValidateGasMeterPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GasMeterPoints); i++ {

		if m.GasMeterPoints[i] != nil {

			if swag.IsZero(m.GasMeterPoints[i]) { // not required
				return nil
			}

			if err := m.GasMeterPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gas_meter_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gas_meter_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
