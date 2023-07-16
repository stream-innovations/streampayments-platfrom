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

// DeleteWalletReader is a Reader for the DeleteWallet structure.
type DeleteWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWalletNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWalletNoContent creates a DeleteWalletNoContent with default headers values
func NewDeleteWalletNoContent() *DeleteWalletNoContent {
	return &DeleteWalletNoContent{}
}

/* DeleteWalletNoContent describes a response with status code 204, with default header values.

Wallet deleted
*/
type DeleteWalletNoContent struct {
}

func (o *DeleteWalletNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wallet/{walletId}][%d] deleteWalletNoContent ", 204)
}

func (o *DeleteWalletNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWalletBadRequest creates a DeleteWalletBadRequest with default headers values
func NewDeleteWalletBadRequest() *DeleteWalletBadRequest {
	return &DeleteWalletBadRequest{}
}

/* DeleteWalletBadRequest describes a response with status code 400, with default header values.

Validation error / Not found
*/
type DeleteWalletBadRequest struct {
	Payload *model.ErrorResponse
}

func (o *DeleteWalletBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /wallet/{walletId}][%d] deleteWalletBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteWalletBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *DeleteWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
