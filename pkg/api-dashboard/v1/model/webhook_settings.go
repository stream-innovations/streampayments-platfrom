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

// WebhookSettings webhook settings
//
// swagger:model webhookSettings
type WebhookSettings struct {

	// HMAC secret for checking webhook signature
	// Example: xa9iZoo6
	// Max Length: 128
	Secret string `json:"secret"`

	// Webhook URL
	// Example: https://my-site.com/webhook/streampay-pay
	// Required: true
	URL string `json:"url"`
}

// Validate validates this webhook settings
func (m *WebhookSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookSettings) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if err := validate.MaxLength("secret", "body", m.Secret, 128); err != nil {
		return err
	}

	return nil
}

func (m *WebhookSettings) validateURL(formats strfmt.Registry) error {

	if err := validate.RequiredString("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this webhook settings based on context it is used
func (m *WebhookSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookSettings) UnmarshalBinary(b []byte) error {
	var res WebhookSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
