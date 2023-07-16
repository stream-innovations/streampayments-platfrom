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

// CurrencyExchangeRate currency exchange rate
//
// swagger:model currencyExchangeRate
type CurrencyExchangeRate struct {

	// crypto amount
	// Example: 0.1231232453453
	// Required: true
	CryptoAmount string `json:"cryptoAmount"`

	// crypto currency
	// Example: DAI
	// Required: true
	CryptoCurrency string `json:"cryptoCurrency"`

	// display name
	// Example: USD → ETH_DAI
	// Required: true
	DisplayName string `json:"displayName"`

	// exchange rate
	// Example: 51.1
	// Required: true
	ExchangeRate float64 `json:"exchangeRate"`

	// fiat amount
	// Example: 49.9
	// Required: true
	// Minimum: 0.01
	FiatAmount float64 `json:"fiatAmount"`

	// fiat currency
	// Example: USD
	// Required: true
	FiatCurrency string `json:"fiatCurrency"`

	// network
	// Example: ETH
	// Required: true
	Network string `json:"network"`
}

// Validate validates this currency exchange rate
func (m *CurrencyExchangeRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCryptoAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCryptoCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiatAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiatCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrencyExchangeRate) validateCryptoAmount(formats strfmt.Registry) error {

	if err := validate.RequiredString("cryptoAmount", "body", m.CryptoAmount); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateCryptoCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("cryptoCurrency", "body", m.CryptoCurrency); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.RequiredString("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateExchangeRate(formats strfmt.Registry) error {

	if err := validate.Required("exchangeRate", "body", float64(m.ExchangeRate)); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateFiatAmount(formats strfmt.Registry) error {

	if err := validate.Required("fiatAmount", "body", float64(m.FiatAmount)); err != nil {
		return err
	}

	if err := validate.Minimum("fiatAmount", "body", m.FiatAmount, 0.01, false); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateFiatCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("fiatCurrency", "body", m.FiatCurrency); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyExchangeRate) validateNetwork(formats strfmt.Registry) error {

	if err := validate.RequiredString("network", "body", m.Network); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this currency exchange rate based on context it is used
func (m *CurrencyExchangeRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyExchangeRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyExchangeRate) UnmarshalBinary(b []byte) error {
	var res CurrencyExchangeRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
