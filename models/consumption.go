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

// Consumption consumption
//
// swagger:model Consumption
type Consumption struct {

	// consumption
	// Required: true
	// Read Only: true
	Consumption float64 `json:"consumption"`

	// interval end
	// Required: true
	// Format: date-time
	IntervalEnd *strfmt.DateTime `json:"interval_end"`

	// interval start
	// Required: true
	// Format: date-time
	IntervalStart *strfmt.DateTime `json:"interval_start"`
}

// Validate validates this consumption
func (m *Consumption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Consumption) validateConsumption(formats strfmt.Registry) error {

	if err := validate.Required("consumption", "body", float64(m.Consumption)); err != nil {
		return err
	}

	return nil
}

func (m *Consumption) validateIntervalEnd(formats strfmt.Registry) error {

	if err := validate.Required("interval_end", "body", m.IntervalEnd); err != nil {
		return err
	}

	if err := validate.FormatOf("interval_end", "body", "date-time", m.IntervalEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Consumption) validateIntervalStart(formats strfmt.Registry) error {

	if err := validate.Required("interval_start", "body", m.IntervalStart); err != nil {
		return err
	}

	if err := validate.FormatOf("interval_start", "body", "date-time", m.IntervalStart.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consumption based on the context it is used
func (m *Consumption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsumption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Consumption) contextValidateConsumption(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "consumption", "body", float64(m.Consumption)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Consumption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Consumption) UnmarshalBinary(b []byte) error {
	var res Consumption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
