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

// HistoricalCharge historical charge
//
// swagger:model HistoricalCharge
type HistoricalCharge struct {

	// payment method
	// Required: true
	// Read Only: true
	PaymentMethod *string `json:"payment_method"`

	// valid from
	// Required: true
	// Read Only: true
	// Format: date-time
	ValidFrom *strfmt.DateTime `json:"valid_from"`

	// valid to
	// Required: true
	// Read Only: true
	// Format: date-time
	ValidTo *strfmt.DateTime `json:"valid_to"`

	// value exc vat
	// Required: true
	// Read Only: true
	ValueExcVat float64 `json:"value_exc_vat"`

	// value inc vat
	// Required: true
	// Read Only: true
	ValueIncVat float64 `json:"value_inc_vat"`
}

// Validate validates this historical charge
func (m *HistoricalCharge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueExcVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueIncVat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoricalCharge) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("payment_method", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("valid_from", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("valid_from", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) validateValidTo(formats strfmt.Registry) error {

	if err := validate.Required("valid_to", "body", m.ValidTo); err != nil {
		return err
	}

	if err := validate.FormatOf("valid_to", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) validateValueExcVat(formats strfmt.Registry) error {

	if err := validate.Required("value_exc_vat", "body", float64(m.ValueExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) validateValueIncVat(formats strfmt.Registry) error {

	if err := validate.Required("value_inc_vat", "body", float64(m.ValueIncVat)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this historical charge based on the context it is used
func (m *HistoricalCharge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePaymentMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValueExcVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValueIncVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoricalCharge) contextValidatePaymentMethod(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "payment_method", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) contextValidateValidFrom(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "valid_from", "body", m.ValidFrom); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) contextValidateValidTo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "valid_to", "body", m.ValidTo); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) contextValidateValueExcVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value_exc_vat", "body", float64(m.ValueExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *HistoricalCharge) contextValidateValueIncVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value_inc_vat", "body", float64(m.ValueIncVat)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoricalCharge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoricalCharge) UnmarshalBinary(b []byte) error {
	var res HistoricalCharge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
