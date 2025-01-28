// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BespokePPSTariffRates bespoke p p s tariff rates
//
// swagger:model BespokePPSTariffRates
type BespokePPSTariffRates struct {

	// The payment method for the rate.
	//
	// * `DD` - Direct Debit
	// * `NDD` - Non-Direct Debit
	// * `PP` - Prepayment
	// Required: true
	// Enum: ["DD","NDD","PP"]
	PaymentMethod *string `json:"payment_method"`

	// The value in pence per day of the charge (excluding VAT).
	// Pattern: ^-?\d{0,7}(?:\.\d{0,5})?$
	StandingCharge string `json:"standing_charge,omitempty"`

	// The value in pence per kWh of the charge (excluding VAT).
	//
	// This field should be used for gas meters.
	// Pattern: ^-?\d{0,7}(?:\.\d{0,5})?$
	UnitRate string `json:"unit_rate,omitempty"`

	// List of the value in pence per kWh of the charges (excluding VAT).
	//
	// For elec meters, the unit rates are provided on a per register basis in this array.
	UnitRates []*BespokeElectricityUnitRate `json:"unit_rates"`
}

// Validate validates this bespoke p p s tariff rates
func (m *BespokePPSTariffRates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingCharge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitRates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bespokePPSTariffRatesTypePaymentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DD","NDD","PP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bespokePPSTariffRatesTypePaymentMethodPropEnum = append(bespokePPSTariffRatesTypePaymentMethodPropEnum, v)
	}
}

const (

	// BespokePPSTariffRatesPaymentMethodDD captures enum value "DD"
	BespokePPSTariffRatesPaymentMethodDD string = "DD"

	// BespokePPSTariffRatesPaymentMethodNDD captures enum value "NDD"
	BespokePPSTariffRatesPaymentMethodNDD string = "NDD"

	// BespokePPSTariffRatesPaymentMethodPP captures enum value "PP"
	BespokePPSTariffRatesPaymentMethodPP string = "PP"
)

// prop value enum
func (m *BespokePPSTariffRates) validatePaymentMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bespokePPSTariffRatesTypePaymentMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BespokePPSTariffRates) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("payment_method", "body", m.PaymentMethod); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentMethodEnum("payment_method", "body", *m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *BespokePPSTariffRates) validateStandingCharge(formats strfmt.Registry) error {
	if swag.IsZero(m.StandingCharge) { // not required
		return nil
	}

	if err := validate.Pattern("standing_charge", "body", m.StandingCharge, `^-?\d{0,7}(?:\.\d{0,5})?$`); err != nil {
		return err
	}

	return nil
}

func (m *BespokePPSTariffRates) validateUnitRate(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitRate) { // not required
		return nil
	}

	if err := validate.Pattern("unit_rate", "body", m.UnitRate, `^-?\d{0,7}(?:\.\d{0,5})?$`); err != nil {
		return err
	}

	return nil
}

func (m *BespokePPSTariffRates) validateUnitRates(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitRates) { // not required
		return nil
	}

	for i := 0; i < len(m.UnitRates); i++ {
		if swag.IsZero(m.UnitRates[i]) { // not required
			continue
		}

		if m.UnitRates[i] != nil {
			if err := m.UnitRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unit_rates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unit_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bespoke p p s tariff rates based on the context it is used
func (m *BespokePPSTariffRates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnitRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BespokePPSTariffRates) contextValidateUnitRates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UnitRates); i++ {

		if m.UnitRates[i] != nil {

			if swag.IsZero(m.UnitRates[i]) { // not required
				return nil
			}

			if err := m.UnitRates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unit_rates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("unit_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BespokePPSTariffRates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BespokePPSTariffRates) UnmarshalBinary(b []byte) error {
	var res BespokePPSTariffRates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
