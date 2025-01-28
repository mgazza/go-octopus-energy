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

// StandardElectricityTariff standard electricity tariff
//
// swagger:model StandardElectricityTariff
type StandardElectricityTariff struct {

	// code
	// Required: true
	// Max Length: 255
	Code *string `json:"code"`

	// dual fuel discount exc vat
	// Required: true
	// Read Only: true
	DualFuelDiscountExcVat float64 `json:"dual_fuel_discount_exc_vat"`

	// dual fuel discount inc vat
	// Required: true
	// Read Only: true
	DualFuelDiscountIncVat float64 `json:"dual_fuel_discount_inc_vat"`

	// exit fees exc vat
	// Required: true
	// Read Only: true
	ExitFeesExcVat int64 `json:"exit_fees_exc_vat"`

	// exit fees inc vat
	// Required: true
	// Read Only: true
	ExitFeesIncVat int64 `json:"exit_fees_inc_vat"`

	// exit fees type
	// Required: true
	// Read Only: true
	ExitFeesType string `json:"exit_fees_type"`

	// links
	// Required: true
	// Read Only: true
	Links []*Link `json:"links"`

	// online discount exc vat
	// Required: true
	// Read Only: true
	OnlineDiscountExcVat float64 `json:"online_discount_exc_vat"`

	// online discount inc vat
	// Required: true
	// Read Only: true
	OnlineDiscountIncVat float64 `json:"online_discount_inc_vat"`

	// standard unit rate exc vat
	// Required: true
	// Read Only: true
	StandardUnitRateExcVat *float64 `json:"standard_unit_rate_exc_vat"`

	// standard unit rate inc vat
	// Required: true
	// Read Only: true
	StandardUnitRateIncVat *float64 `json:"standard_unit_rate_inc_vat"`

	// standing charge exc vat
	// Required: true
	// Read Only: true
	StandingChargeExcVat *float64 `json:"standing_charge_exc_vat"`

	// standing charge inc vat
	// Required: true
	// Read Only: true
	StandingChargeIncVat *float64 `json:"standing_charge_inc_vat"`
}

// Validate validates this standard electricity tariff
func (m *StandardElectricityTariff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualFuelDiscountExcVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDualFuelDiscountIncVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitFeesExcVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitFeesIncVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitFeesType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnlineDiscountExcVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnlineDiscountIncVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardUnitRateExcVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandardUnitRateIncVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingChargeExcVat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingChargeIncVat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardElectricityTariff) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if err := validate.MaxLength("code", "body", *m.Code, 255); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateDualFuelDiscountExcVat(formats strfmt.Registry) error {

	if err := validate.Required("dual_fuel_discount_exc_vat", "body", float64(m.DualFuelDiscountExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateDualFuelDiscountIncVat(formats strfmt.Registry) error {

	if err := validate.Required("dual_fuel_discount_inc_vat", "body", float64(m.DualFuelDiscountIncVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateExitFeesExcVat(formats strfmt.Registry) error {

	if err := validate.Required("exit_fees_exc_vat", "body", int64(m.ExitFeesExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateExitFeesIncVat(formats strfmt.Registry) error {

	if err := validate.Required("exit_fees_inc_vat", "body", int64(m.ExitFeesIncVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateExitFeesType(formats strfmt.Registry) error {

	if err := validate.RequiredString("exit_fees_type", "body", m.ExitFeesType); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StandardElectricityTariff) validateOnlineDiscountExcVat(formats strfmt.Registry) error {

	if err := validate.Required("online_discount_exc_vat", "body", float64(m.OnlineDiscountExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateOnlineDiscountIncVat(formats strfmt.Registry) error {

	if err := validate.Required("online_discount_inc_vat", "body", float64(m.OnlineDiscountIncVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateStandardUnitRateExcVat(formats strfmt.Registry) error {

	if err := validate.Required("standard_unit_rate_exc_vat", "body", m.StandardUnitRateExcVat); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateStandardUnitRateIncVat(formats strfmt.Registry) error {

	if err := validate.Required("standard_unit_rate_inc_vat", "body", m.StandardUnitRateIncVat); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateStandingChargeExcVat(formats strfmt.Registry) error {

	if err := validate.Required("standing_charge_exc_vat", "body", m.StandingChargeExcVat); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) validateStandingChargeIncVat(formats strfmt.Registry) error {

	if err := validate.Required("standing_charge_inc_vat", "body", m.StandingChargeIncVat); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this standard electricity tariff based on the context it is used
func (m *StandardElectricityTariff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDualFuelDiscountExcVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDualFuelDiscountIncVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExitFeesExcVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExitFeesIncVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExitFeesType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnlineDiscountExcVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnlineDiscountIncVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardUnitRateExcVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandardUnitRateIncVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandingChargeExcVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandingChargeIncVat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardElectricityTariff) contextValidateDualFuelDiscountExcVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dual_fuel_discount_exc_vat", "body", float64(m.DualFuelDiscountExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateDualFuelDiscountIncVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dual_fuel_discount_inc_vat", "body", float64(m.DualFuelDiscountIncVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateExitFeesExcVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exit_fees_exc_vat", "body", int64(m.ExitFeesExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateExitFeesIncVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exit_fees_inc_vat", "body", int64(m.ExitFeesIncVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateExitFeesType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exit_fees_type", "body", string(m.ExitFeesType)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "links", "body", []*Link(m.Links)); err != nil {
		return err
	}

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {

			if swag.IsZero(m.Links[i]) { // not required
				return nil
			}

			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateOnlineDiscountExcVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "online_discount_exc_vat", "body", float64(m.OnlineDiscountExcVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateOnlineDiscountIncVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "online_discount_inc_vat", "body", float64(m.OnlineDiscountIncVat)); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateStandardUnitRateExcVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "standard_unit_rate_exc_vat", "body", m.StandardUnitRateExcVat); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateStandardUnitRateIncVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "standard_unit_rate_inc_vat", "body", m.StandardUnitRateIncVat); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateStandingChargeExcVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "standing_charge_exc_vat", "body", m.StandingChargeExcVat); err != nil {
		return err
	}

	return nil
}

func (m *StandardElectricityTariff) contextValidateStandingChargeIncVat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "standing_charge_inc_vat", "body", m.StandingChargeIncVat); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandardElectricityTariff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandardElectricityTariff) UnmarshalBinary(b []byte) error {
	var res StandardElectricityTariff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
