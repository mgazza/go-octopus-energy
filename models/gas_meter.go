// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GasMeter gas meter
//
// swagger:model GasMeter
type GasMeter struct {

	// serial number
	// Example: 12345678901234
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this gas meter
func (m *GasMeter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gas meter based on context it is used
func (m *GasMeter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GasMeter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GasMeter) UnmarshalBinary(b []byte) error {
	var res GasMeter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
