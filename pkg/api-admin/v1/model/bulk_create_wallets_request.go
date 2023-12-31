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

// BulkCreateWalletsRequest bulk create wallets request
//
// swagger:model bulkCreateWalletsRequest
type BulkCreateWalletsRequest struct {

	// amount
	// Required: true
	// Maximum: 32
	// Minimum: 1
	Amount int64 `json:"amount"`

	// blockchain
	// Required: true
	Blockchain *Blockchain `json:"blockchain"`
}

// Validate validates this bulk create wallets request
func (m *BulkCreateWalletsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockchain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCreateWalletsRequest) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", int64(m.Amount)); err != nil {
		return err
	}

	if err := validate.MinimumInt("amount", "body", m.Amount, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amount", "body", m.Amount, 32, false); err != nil {
		return err
	}

	return nil
}

func (m *BulkCreateWalletsRequest) validateBlockchain(formats strfmt.Registry) error {

	if err := validate.Required("blockchain", "body", m.Blockchain); err != nil {
		return err
	}

	if err := validate.Required("blockchain", "body", m.Blockchain); err != nil {
		return err
	}

	if m.Blockchain != nil {
		if err := m.Blockchain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blockchain")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bulk create wallets request based on the context it is used
func (m *BulkCreateWalletsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockchain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkCreateWalletsRequest) contextValidateBlockchain(ctx context.Context, formats strfmt.Registry) error {

	if m.Blockchain != nil {
		if err := m.Blockchain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blockchain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkCreateWalletsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkCreateWalletsRequest) UnmarshalBinary(b []byte) error {
	var res BulkCreateWalletsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
