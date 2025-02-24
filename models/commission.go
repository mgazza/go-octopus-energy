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

// Commission commission
//
// swagger:model _Commission
type Commission struct {

	// Sub-domain of the affiliate link due the commission for this agreement.
	// Max Length: 64
	AffiliateLinkSubdomain string `json:"affiliate_link_subdomain,omitempty"`

	// Name of the affiliate organisation due the commission for this agreement.
	// Required: true
	// Max Length: 128
	AffiliateOrganisationName *string `json:"affiliate_organisation_name"`

	// fixed tpi fee
	// Minimum: 0
	FixedTpiFee *int64 `json:"fixed_tpi_fee,omitempty"`

	// Uplift in pence to be paid to the affiliate.
	// Pattern: ^-?\d{0,11}(?:\.\d{0,2})?$
	StandingChargeUplift string `json:"standing_charge_uplift,omitempty"`

	// Uplift in pence to be paid to the affiliate.
	// Required: true
	// Pattern: ^-?\d{0,11}(?:\.\d{0,2})?$
	UnitRateUplift *string `json:"unit_rate_uplift"`
}

// Validate validates this commission
func (m *Commission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffiliateLinkSubdomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAffiliateOrganisationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedTpiFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingChargeUplift(formats); err != nil {
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

func (m *Commission) validateAffiliateLinkSubdomain(formats strfmt.Registry) error {
	if swag.IsZero(m.AffiliateLinkSubdomain) { // not required
		return nil
	}

	if err := validate.MaxLength("affiliate_link_subdomain", "body", m.AffiliateLinkSubdomain, 64); err != nil {
		return err
	}

	return nil
}

func (m *Commission) validateAffiliateOrganisationName(formats strfmt.Registry) error {

	if err := validate.Required("affiliate_organisation_name", "body", m.AffiliateOrganisationName); err != nil {
		return err
	}

	if err := validate.MaxLength("affiliate_organisation_name", "body", *m.AffiliateOrganisationName, 128); err != nil {
		return err
	}

	return nil
}

func (m *Commission) validateFixedTpiFee(formats strfmt.Registry) error {
	if swag.IsZero(m.FixedTpiFee) { // not required
		return nil
	}

	if err := validate.MinimumInt("fixed_tpi_fee", "body", *m.FixedTpiFee, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Commission) validateStandingChargeUplift(formats strfmt.Registry) error {
	if swag.IsZero(m.StandingChargeUplift) { // not required
		return nil
	}

	if err := validate.Pattern("standing_charge_uplift", "body", m.StandingChargeUplift, `^-?\d{0,11}(?:\.\d{0,2})?$`); err != nil {
		return err
	}

	return nil
}

func (m *Commission) validateUnitRateUplift(formats strfmt.Registry) error {

	if err := validate.Required("unit_rate_uplift", "body", m.UnitRateUplift); err != nil {
		return err
	}

	if err := validate.Pattern("unit_rate_uplift", "body", *m.UnitRateUplift, `^-?\d{0,11}(?:\.\d{0,2})?$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this commission based on context it is used
func (m *Commission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Commission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commission) UnmarshalBinary(b []byte) error {
	var res Commission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
