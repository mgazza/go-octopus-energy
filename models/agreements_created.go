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

// AgreementsCreated agreements created
//
// swagger:model AgreementsCreated
type AgreementsCreated struct {

	// account number
	// Required: true
	AccountNumber *string `json:"account_number"`

	// agreements
	// Required: true
	Agreements []*Agreement `json:"agreements"`
}

// Validate validates this agreements created
func (m *AgreementsCreated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgreements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgreementsCreated) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("account_number", "body", m.AccountNumber); err != nil {
		return err
	}

	return nil
}

func (m *AgreementsCreated) validateAgreements(formats strfmt.Registry) error {

	if err := validate.Required("agreements", "body", m.Agreements); err != nil {
		return err
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

// ContextValidate validate this agreements created based on the context it is used
func (m *AgreementsCreated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgreements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgreementsCreated) contextValidateAgreements(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AgreementsCreated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgreementsCreated) UnmarshalBinary(b []byte) error {
	var res AgreementsCreated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
