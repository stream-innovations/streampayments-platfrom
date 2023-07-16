// Code generated by go-swagger; DO NOT EDIT.

package model

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

// Conversion Exchange rate
//
// swagger:model conversion
type Conversion struct {

	// Converted amount
	// Example: 0.066
	ConvertedAmount string `json:"convertedAmount,omitempty"`

	// Conversion rate
	// Example: 1820.5
	ExchangeRate float64 `json:"exchangeRate,omitempty"`

	// Selected ticker
	// Example: USD
	From string `json:"from,omitempty"`

	// Selected ticker type
	// Enum: [fiat crypto]
	FromType string `json:"fromType,omitempty"`

	// Selected amount
	// Example: 123
	SelectedAmount string `json:"selectedAmount,omitempty"`

	// Desired ticket
	// Example: ETH
	To string `json:"to,omitempty"`

	// Desired ticker type
	// Enum: [fiat crypto]
	ToType string `json:"toType,omitempty"`
}

// Validate validates this conversion
func (m *Conversion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversionTypeFromTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fiat","crypto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversionTypeFromTypePropEnum = append(conversionTypeFromTypePropEnum, v)
	}
}

const (

	// ConversionFromTypeFiat captures enum value "fiat"
	ConversionFromTypeFiat string = "fiat"

	// ConversionFromTypeCrypto captures enum value "crypto"
	ConversionFromTypeCrypto string = "crypto"
)

// prop value enum
func (m *Conversion) validateFromTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversionTypeFromTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Conversion) validateFromType(formats strfmt.Registry) error {
	if swag.IsZero(m.FromType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFromTypeEnum("fromType", "body", m.FromType); err != nil {
		return err
	}

	return nil
}

var conversionTypeToTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fiat","crypto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversionTypeToTypePropEnum = append(conversionTypeToTypePropEnum, v)
	}
}

const (

	// ConversionToTypeFiat captures enum value "fiat"
	ConversionToTypeFiat string = "fiat"

	// ConversionToTypeCrypto captures enum value "crypto"
	ConversionToTypeCrypto string = "crypto"
)

// prop value enum
func (m *Conversion) validateToTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversionTypeToTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Conversion) validateToType(formats strfmt.Registry) error {
	if swag.IsZero(m.ToType) { // not required
		return nil
	}

	// value enum
	if err := m.validateToTypeEnum("toType", "body", m.ToType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conversion based on context it is used
func (m *Conversion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Conversion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Conversion) UnmarshalBinary(b []byte) error {
	var res Conversion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}