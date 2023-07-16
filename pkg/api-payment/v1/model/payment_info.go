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

// PaymentInfo Payment info
//
// swagger:model paymentInfo
type PaymentInfo struct {

	// Amount to pay in selected currency. Wei for ETH, Satoshi for BTC and so on...
	// Example: 1000000
	// Required: true
	Amount string `json:"amount"`

	// Human readable amount to pay.
	// Example: 0.123
	// Required: true
	AmountFormatted string `json:"amountFormatted"`

	// Expiration duration in minutes
	// Example: 20
	// Required: true
	ExpirationDurationMin int64 `json:"expirationDurationMin"`

	// payment expiration timestamp (UTC)
	// Example: 2023-03-09T20:18:07.809Z
	// Required: true
	// Format: datetime
	ExpiresAt strfmt.DateTime `json:"expiresAt"`

	// Payment link for QR code
	// Example: tron:TVEaDaTKJZ2RsQUWREWykouuHak9scyZaf
	// Required: true
	PaymentLink string `json:"paymentLink"`

	// recipient address
	// Example: 0xbca4a8417e823484b21d2f4f1f1324d951236a49
	// Required: true
	RecipientAddress string `json:"recipientAddress"`

	// Payment status
	// Example: success
	// Required: true
	// Enum: [pending inProgress success]
	Status string `json:"status"`

	// Success action. Present if payment is successful
	// Enum: [redirect showMessage]
	SuccessAction *string `json:"successAction,omitempty"`

	// Success message visible after payment confirmation (if successAction == `showMessage`)
	//
	// Example: Thanks you for purchasing our product! You will receive invitation on your email\n
	SuccessMessage *string `json:"successMessage,omitempty"`

	// Success URL to where a user should be redirected after payment confirmation.
	// Empty if success action is `showMessage`
	//
	// Example: https://site.com/success
	SuccessURL *string `json:"successUrl,omitempty"`
}

// Validate validates this payment info
func (m *PaymentInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountFormatted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDurationMin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentInfo) validateAmount(formats strfmt.Registry) error {

	if err := validate.RequiredString("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *PaymentInfo) validateAmountFormatted(formats strfmt.Registry) error {

	if err := validate.RequiredString("amountFormatted", "body", m.AmountFormatted); err != nil {
		return err
	}

	return nil
}

func (m *PaymentInfo) validateExpirationDurationMin(formats strfmt.Registry) error {

	if err := validate.Required("expirationDurationMin", "body", int64(m.ExpirationDurationMin)); err != nil {
		return err
	}

	return nil
}

func (m *PaymentInfo) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expiresAt", "body", strfmt.DateTime(m.ExpiresAt)); err != nil {
		return err
	}

	if err := validate.FormatOf("expiresAt", "body", "datetime", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentInfo) validatePaymentLink(formats strfmt.Registry) error {

	if err := validate.RequiredString("paymentLink", "body", m.PaymentLink); err != nil {
		return err
	}

	return nil
}

func (m *PaymentInfo) validateRecipientAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("recipientAddress", "body", m.RecipientAddress); err != nil {
		return err
	}

	return nil
}

var paymentInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","inProgress","success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentInfoTypeStatusPropEnum = append(paymentInfoTypeStatusPropEnum, v)
	}
}

const (

	// PaymentInfoStatusPending captures enum value "pending"
	PaymentInfoStatusPending string = "pending"

	// PaymentInfoStatusInProgress captures enum value "inProgress"
	PaymentInfoStatusInProgress string = "inProgress"

	// PaymentInfoStatusSuccess captures enum value "success"
	PaymentInfoStatusSuccess string = "success"
)

// prop value enum
func (m *PaymentInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var paymentInfoTypeSuccessActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["redirect","showMessage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentInfoTypeSuccessActionPropEnum = append(paymentInfoTypeSuccessActionPropEnum, v)
	}
}

const (

	// PaymentInfoSuccessActionRedirect captures enum value "redirect"
	PaymentInfoSuccessActionRedirect string = "redirect"

	// PaymentInfoSuccessActionShowMessage captures enum value "showMessage"
	PaymentInfoSuccessActionShowMessage string = "showMessage"
)

// prop value enum
func (m *PaymentInfo) validateSuccessActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, paymentInfoTypeSuccessActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PaymentInfo) validateSuccessAction(formats strfmt.Registry) error {
	if swag.IsZero(m.SuccessAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateSuccessActionEnum("successAction", "body", *m.SuccessAction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this payment info based on context it is used
func (m *PaymentInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentInfo) UnmarshalBinary(b []byte) error {
	var res PaymentInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
