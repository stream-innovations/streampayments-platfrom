// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"https://github.com/stream-innovations/streampayments-platfrom/pkg/api-kms/v1/model"
)

// CreateTronTransactionReader is a Reader for the CreateTronTransaction structure.
type CreateTronTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTronTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTronTransactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTronTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTronTransactionCreated creates a CreateTronTransactionCreated with default headers values
func NewCreateTronTransactionCreated() *CreateTronTransactionCreated {
	return &CreateTronTransactionCreated{}
}

/* CreateTronTransactionCreated describes a response with status code 201, with default header values.

Transaction Created
*/
type CreateTronTransactionCreated struct {
	Payload *model.TronTransaction
}

func (o *CreateTronTransactionCreated) Error() string {
	return fmt.Sprintf("[POST /wallet/{walletId}/transaction/tron][%d] createTronTransactionCreated  %+v", 201, o.Payload)
}
func (o *CreateTronTransactionCreated) GetPayload() *model.TronTransaction {
	return o.Payload
}

func (o *CreateTronTransactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.TronTransaction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTronTransactionBadRequest creates a CreateTronTransactionBadRequest with default headers values
func NewCreateTronTransactionBadRequest() *CreateTronTransactionBadRequest {
	return &CreateTronTransactionBadRequest{}
}

/* CreateTronTransactionBadRequest describes a response with status code 400, with default header values.

Validation error / Not found
*/
type CreateTronTransactionBadRequest struct {
	Payload *model.ErrorResponse
}

func (o *CreateTronTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /wallet/{walletId}/transaction/tron][%d] createTronTransactionBadRequest  %+v", 400, o.Payload)
}
func (o *CreateTronTransactionBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *CreateTronTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
