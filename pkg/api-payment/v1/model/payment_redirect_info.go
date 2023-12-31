// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentRedirectInfo payment redirect info
//
// swagger:model paymentRedirectInfo
type PaymentRedirectInfo struct {

	// Payment UUID
	// Example: a51e7a5-f0c8-48dc-a9fb-a335481ae846
	// Required: true
	ID string `json:"id"`
}

// Validate validates this payment redirect info
func (m *PaymentRedirectInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentRedirectInfo) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this payment redirect info based on context it is used
func (m *PaymentRedirectInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRedirectInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRedirectInfo) UnmarshalBinary(b []byte) error {
	var res PaymentRedirectInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
