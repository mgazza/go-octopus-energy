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

// ElectricityMeterPoint electricity meter point
//
// swagger:model ElectricityMeterPoint
type ElectricityMeterPoint struct {

	// address
	Address *Address `json:"address,omitempty"`

	// If applicable, must contain two objects: one for direct debit bespoke and another for non-direct debit bespoke rates. Cannot be used in conjunction with 'bespoke_tariff_rates'.
	// Max Items: 2
	// Min Items: 0
	BespokePpsTariffRates []*BespokePPSTariffRates `json:"bespoke_pps_tariff_rates"`

	// A json object containing bespoke tariff rates, if applicable.
	//
	// Cannot be used in conjunction with 'bespoke_pps_tariff_rates'.
	BespokeTariffRates struct {
		BespokeTariffRates
	} `json:"bespoke_tariff_rates,omitempty"`

	// The expected annual consumption for the day/peak register in kWh.
	//
	// Required if eco7 or three-rate tariff.
	// Maximum: 500000
	// Minimum: 1
	ConsumptionDay int64 `json:"consumption_day,omitempty"`

	// The expected annual consumption for the night/off-peak register in kWh.
	//
	// Required if eco7 or three-rate tariff.
	// Maximum: 500000
	// Minimum: 0
	ConsumptionNight *int64 `json:"consumption_night,omitempty"`

	// The expected annual consumption for the additional off-peak register in kWh.
	//
	// Required if three-rate tariff.
	// Maximum: 500000
	// Minimum: 0
	ConsumptionOffPeak *int64 `json:"consumption_off_peak,omitempty"`

	// The expected annual consumption for the standard register in kWh.
	//
	// Required if standard tariff.
	// Maximum: 500000
	// Minimum: 1
	ConsumptionStandard int64 `json:"consumption_standard,omitempty"`

	// current supplier name
	// Max Length: 255
	CurrentSupplierName string `json:"current_supplier_name,omitempty"`

	// current supplier tariff
	// Max Length: 255
	CurrentSupplierTariff string `json:"current_supplier_tariff,omitempty"`

	// fixed tpi fee
	// Minimum: 0
	FixedTpiFee *int64 `json:"fixed_tpi_fee,omitempty"`

	// Whether a smart meter is installed. The default is 'false'.
	HasSmartMeter bool `json:"has_smart_meter,omitempty"`

	// The default is 'CREDIT', which represents any kind of traditional credit meter or smart meter. This field is mandatory for traditional prepay meters, for which the value 'PREPAYMENT' should be used.
	//
	// * `CREDIT` - CREDIT
	// * `PREPAYMENT` - PREPAYMENT
	// Enum: ["CREDIT","PREPAYMENT"]
	MeterType *string `json:"meter_type,omitempty"`

	// The meter-point identifier.
	// Max Length: 13
	Mpan string `json:"mpan,omitempty"`

	// preferred ssd
	// Format: date
	PreferredSsd strfmt.Date `json:"preferred_ssd,omitempty"`

	// quote
	// Required: true
	Quote *Quote `json:"quote"`

	// The ID of the market supply quoted product returned by the 'QuoteNewMeterPoints' or 'QuoteNewMeterPointsOnBespokeProducts' mutation which was selected by the customer for this meter point. Required if a 'quote_request_code' is supplied.
	QuotedProductID int64 `json:"quoted_product_id,omitempty"`

	// The amount to be added to the standing charge as commission (a string, pence per day).
	// Pattern: ^-?\d{0,4}(?:\.\d{0,2})?$
	StandingChargeUplift string `json:"standing_charge_uplift,omitempty"`

	// The tariff code as stored by the supplier.
	//
	// Must match an available tariff on the 'sold_at' datetime for the authenticating partner. Required if no 'quote_request_code' is supplied.
	// Max Length: 255
	TariffCode string `json:"tariff_code,omitempty"`

	// The amount to be added to the unit rate as commission (a string, pence per kWh).
	// Pattern: ^-?\d{0,2}(?:\.\d{0,2})?$
	UnitRateUplift string `json:"unit_rate_uplift,omitempty"`
}

// Validate validates this electricity meter point
func (m *ElectricityMeterPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBespokePpsTariffRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBespokeTariffRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumptionDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumptionNight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumptionOffPeak(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsumptionStandard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentSupplierName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentSupplierTariff(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedTpiFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMpan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredSsd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingChargeUplift(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTariffCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitRateUplift(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElectricityMeterPoint) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *ElectricityMeterPoint) validateBespokePpsTariffRates(formats strfmt.Registry) error {
	if swag.IsZero(m.BespokePpsTariffRates) { // not required
		return nil
	}

	iBespokePpsTariffRatesSize := int64(len(m.BespokePpsTariffRates))

	if err := validate.MinItems("bespoke_pps_tariff_rates", "body", iBespokePpsTariffRatesSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("bespoke_pps_tariff_rates", "body", iBespokePpsTariffRatesSize, 2); err != nil {
		return err
	}

	for i := 0; i < len(m.BespokePpsTariffRates); i++ {
		if swag.IsZero(m.BespokePpsTariffRates[i]) { // not required
			continue
		}

		if m.BespokePpsTariffRates[i] != nil {
			if err := m.BespokePpsTariffRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bespoke_pps_tariff_rates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bespoke_pps_tariff_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElectricityMeterPoint) validateBespokeTariffRates(formats strfmt.Registry) error {
	if swag.IsZero(m.BespokeTariffRates) { // not required
		return nil
	}

	return nil
}

func (m *ElectricityMeterPoint) validateConsumptionDay(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionDay) { // not required
		return nil
	}

	if err := validate.MinimumInt("consumption_day", "body", m.ConsumptionDay, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("consumption_day", "body", m.ConsumptionDay, 500000, false); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateConsumptionNight(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionNight) { // not required
		return nil
	}

	if err := validate.MinimumInt("consumption_night", "body", *m.ConsumptionNight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("consumption_night", "body", *m.ConsumptionNight, 500000, false); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateConsumptionOffPeak(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionOffPeak) { // not required
		return nil
	}

	if err := validate.MinimumInt("consumption_off_peak", "body", *m.ConsumptionOffPeak, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("consumption_off_peak", "body", *m.ConsumptionOffPeak, 500000, false); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateConsumptionStandard(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsumptionStandard) { // not required
		return nil
	}

	if err := validate.MinimumInt("consumption_standard", "body", m.ConsumptionStandard, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("consumption_standard", "body", m.ConsumptionStandard, 500000, false); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateCurrentSupplierName(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentSupplierName) { // not required
		return nil
	}

	if err := validate.MaxLength("current_supplier_name", "body", m.CurrentSupplierName, 255); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateCurrentSupplierTariff(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentSupplierTariff) { // not required
		return nil
	}

	if err := validate.MaxLength("current_supplier_tariff", "body", m.CurrentSupplierTariff, 255); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateFixedTpiFee(formats strfmt.Registry) error {
	if swag.IsZero(m.FixedTpiFee) { // not required
		return nil
	}

	if err := validate.MinimumInt("fixed_tpi_fee", "body", *m.FixedTpiFee, 0, false); err != nil {
		return err
	}

	return nil
}

var electricityMeterPointTypeMeterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREDIT","PREPAYMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		electricityMeterPointTypeMeterTypePropEnum = append(electricityMeterPointTypeMeterTypePropEnum, v)
	}
}

const (

	// ElectricityMeterPointMeterTypeCREDIT captures enum value "CREDIT"
	ElectricityMeterPointMeterTypeCREDIT string = "CREDIT"

	// ElectricityMeterPointMeterTypePREPAYMENT captures enum value "PREPAYMENT"
	ElectricityMeterPointMeterTypePREPAYMENT string = "PREPAYMENT"
)

// prop value enum
func (m *ElectricityMeterPoint) validateMeterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, electricityMeterPointTypeMeterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ElectricityMeterPoint) validateMeterType(formats strfmt.Registry) error {
	if swag.IsZero(m.MeterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMeterTypeEnum("meter_type", "body", *m.MeterType); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateMpan(formats strfmt.Registry) error {
	if swag.IsZero(m.Mpan) { // not required
		return nil
	}

	if err := validate.MaxLength("mpan", "body", m.Mpan, 13); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validatePreferredSsd(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredSsd) { // not required
		return nil
	}

	if err := validate.FormatOf("preferred_ssd", "body", "date", m.PreferredSsd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateQuote(formats strfmt.Registry) error {

	if err := validate.Required("quote", "body", m.Quote); err != nil {
		return err
	}

	if m.Quote != nil {
		if err := m.Quote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quote")
			}
			return err
		}
	}

	return nil
}

func (m *ElectricityMeterPoint) validateStandingChargeUplift(formats strfmt.Registry) error {
	if swag.IsZero(m.StandingChargeUplift) { // not required
		return nil
	}

	if err := validate.Pattern("standing_charge_uplift", "body", m.StandingChargeUplift, `^-?\d{0,4}(?:\.\d{0,2})?$`); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateTariffCode(formats strfmt.Registry) error {
	if swag.IsZero(m.TariffCode) { // not required
		return nil
	}

	if err := validate.MaxLength("tariff_code", "body", m.TariffCode, 255); err != nil {
		return err
	}

	return nil
}

func (m *ElectricityMeterPoint) validateUnitRateUplift(formats strfmt.Registry) error {
	if swag.IsZero(m.UnitRateUplift) { // not required
		return nil
	}

	if err := validate.Pattern("unit_rate_uplift", "body", m.UnitRateUplift, `^-?\d{0,2}(?:\.\d{0,2})?$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this electricity meter point based on the context it is used
func (m *ElectricityMeterPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBespokePpsTariffRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBespokeTariffRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElectricityMeterPoint) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {

		if swag.IsZero(m.Address) { // not required
			return nil
		}

		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *ElectricityMeterPoint) contextValidateBespokePpsTariffRates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BespokePpsTariffRates); i++ {

		if m.BespokePpsTariffRates[i] != nil {

			if swag.IsZero(m.BespokePpsTariffRates[i]) { // not required
				return nil
			}

			if err := m.BespokePpsTariffRates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bespoke_pps_tariff_rates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bespoke_pps_tariff_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElectricityMeterPoint) contextValidateBespokeTariffRates(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ElectricityMeterPoint) contextValidateQuote(ctx context.Context, formats strfmt.Registry) error {

	if m.Quote != nil {

		if err := m.Quote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElectricityMeterPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElectricityMeterPoint) UnmarshalBinary(b []byte) error {
	var res ElectricityMeterPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
