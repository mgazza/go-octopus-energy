// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BespokeElectricityUnitRate bespoke electricity unit rate
//
// swagger:model BespokeElectricityUnitRate
type BespokeElectricityUnitRate struct {

	// The value in pence per kWh of the charge (excluding VAT).
	//
	// This field should be used for gas meters.
	//
	// * `DD` - DD
	// * `NDD` - NDD
	// * `PP` - PP
	// Enum: ["DD","NDD","PP"]
	PaymentMethod string `json:"payment_method,omitempty"`

	// The rate type of the unit rate.
	//
	// If the register_identifier is not provided, the rate type must be provided instead.
	//
	// * `STANDARD` - Standard rate (pence per kWh)
	// * `ECO7_DAY` - Day (or peak) rate (pence per kWh)
	// * `ECO7_NIGHT` - Night (or off-peak) rate (pence per kWh)
	// * `OFF_PEAK` - Additional off-peak rate for three-rate tariffs (pence per kWh)
	// * `SUMMER_PEAK` - Summer peak rate (pence per kWh) for two-rate tariffs
	// * `SUMMER_OFF_PEAK` - Summer off-peak rate (pence per kWh) for two-rate tariffs
	// * `WINTER_PEAK` - Winter peak rate (pence per kWh) for two-rate tariffs
	// * `WINTER_OFF_PEAK` - Winter off-peak rate (pence per kWh) for two-rate tariffs
	// * `NUCLEAR_RAB` - Nuclear RAB rate (pence per day) for business tariffs
	// * `TNUOS` - TNUoS rate (pence per day) for business tariffs
	// * `CAPACITY_MARKET` - Capacity Market rate (pence per day) for business tariffs
	// * `EV_DEVICE` - Electric vehicle device rate (pence per kWh) for sub-meter billing
	// Enum: ["STANDARD","ECO7_DAY","ECO7_NIGHT","OFF_PEAK","SUMMER_PEAK","SUMMER_OFF_PEAK","WINTER_PEAK","WINTER_OFF_PEAK","NUCLEAR_RAB","TNUOS","CAPACITY_MARKET","EV_DEVICE"]
	RateType string `json:"rate_type,omitempty"`

	// The value in pence per kWh of the charge (excluding VAT).
	//
	// This field should be used for gas meters.
	// Required: true
	// Pattern: ^-?\d{0,7}(?:\.\d{0,5})?$
	UnitRate *string `json:"unit_rate"`
}

// Validate validates this bespoke electricity unit rate
func (m *BespokeElectricityUnitRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bespokeElectricityUnitRateTypePaymentMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DD","NDD","PP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bespokeElectricityUnitRateTypePaymentMethodPropEnum = append(bespokeElectricityUnitRateTypePaymentMethodPropEnum, v)
	}
}

const (

	// BespokeElectricityUnitRatePaymentMethodDD captures enum value "DD"
	BespokeElectricityUnitRatePaymentMethodDD string = "DD"

	// BespokeElectricityUnitRatePaymentMethodNDD captures enum value "NDD"
	BespokeElectricityUnitRatePaymentMethodNDD string = "NDD"

	// BespokeElectricityUnitRatePaymentMethodPP captures enum value "PP"
	BespokeElectricityUnitRatePaymentMethodPP string = "PP"
)

// prop value enum
func (m *BespokeElectricityUnitRate) validatePaymentMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bespokeElectricityUnitRateTypePaymentMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BespokeElectricityUnitRate) validatePaymentMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentMethodEnum("payment_method", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

var bespokeElectricityUnitRateTypeRateTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STANDARD","ECO7_DAY","ECO7_NIGHT","OFF_PEAK","SUMMER_PEAK","SUMMER_OFF_PEAK","WINTER_PEAK","WINTER_OFF_PEAK","NUCLEAR_RAB","TNUOS","CAPACITY_MARKET","EV_DEVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bespokeElectricityUnitRateTypeRateTypePropEnum = append(bespokeElectricityUnitRateTypeRateTypePropEnum, v)
	}
}

const (

	// BespokeElectricityUnitRateRateTypeSTANDARD captures enum value "STANDARD"
	BespokeElectricityUnitRateRateTypeSTANDARD string = "STANDARD"

	// BespokeElectricityUnitRateRateTypeECO7DAY captures enum value "ECO7_DAY"
	BespokeElectricityUnitRateRateTypeECO7DAY string = "ECO7_DAY"

	// BespokeElectricityUnitRateRateTypeECO7NIGHT captures enum value "ECO7_NIGHT"
	BespokeElectricityUnitRateRateTypeECO7NIGHT string = "ECO7_NIGHT"

	// BespokeElectricityUnitRateRateTypeOFFPEAK captures enum value "OFF_PEAK"
	BespokeElectricityUnitRateRateTypeOFFPEAK string = "OFF_PEAK"

	// BespokeElectricityUnitRateRateTypeSUMMERPEAK captures enum value "SUMMER_PEAK"
	BespokeElectricityUnitRateRateTypeSUMMERPEAK string = "SUMMER_PEAK"

	// BespokeElectricityUnitRateRateTypeSUMMEROFFPEAK captures enum value "SUMMER_OFF_PEAK"
	BespokeElectricityUnitRateRateTypeSUMMEROFFPEAK string = "SUMMER_OFF_PEAK"

	// BespokeElectricityUnitRateRateTypeWINTERPEAK captures enum value "WINTER_PEAK"
	BespokeElectricityUnitRateRateTypeWINTERPEAK string = "WINTER_PEAK"

	// BespokeElectricityUnitRateRateTypeWINTEROFFPEAK captures enum value "WINTER_OFF_PEAK"
	BespokeElectricityUnitRateRateTypeWINTEROFFPEAK string = "WINTER_OFF_PEAK"

	// BespokeElectricityUnitRateRateTypeNUCLEARRAB captures enum value "NUCLEAR_RAB"
	BespokeElectricityUnitRateRateTypeNUCLEARRAB string = "NUCLEAR_RAB"

	// BespokeElectricityUnitRateRateTypeTNUOS captures enum value "TNUOS"
	BespokeElectricityUnitRateRateTypeTNUOS string = "TNUOS"

	// BespokeElectricityUnitRateRateTypeCAPACITYMARKET captures enum value "CAPACITY_MARKET"
	BespokeElectricityUnitRateRateTypeCAPACITYMARKET string = "CAPACITY_MARKET"

	// BespokeElectricityUnitRateRateTypeEVDEVICE captures enum value "EV_DEVICE"
	BespokeElectricityUnitRateRateTypeEVDEVICE string = "EV_DEVICE"
)

// prop value enum
func (m *BespokeElectricityUnitRate) validateRateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bespokeElectricityUnitRateTypeRateTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BespokeElectricityUnitRate) validateRateType(formats strfmt.Registry) error {
	if swag.IsZero(m.RateType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRateTypeEnum("rate_type", "body", m.RateType); err != nil {
		return err
	}

	return nil
}

func (m *BespokeElectricityUnitRate) validateUnitRate(formats strfmt.Registry) error {

	if err := validate.Required("unit_rate", "body", m.UnitRate); err != nil {
		return err
	}

	if err := validate.Pattern("unit_rate", "body", *m.UnitRate, `^-?\d{0,7}(?:\.\d{0,5})?$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bespoke electricity unit rate based on context it is used
func (m *BespokeElectricityUnitRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BespokeElectricityUnitRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BespokeElectricityUnitRate) UnmarshalBinary(b []byte) error {
	var res BespokeElectricityUnitRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
