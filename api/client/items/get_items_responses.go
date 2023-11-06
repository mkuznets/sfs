// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mkuznets.com/go/sfs/api/models"
)

// GetItemsReader is a Reader for the GetItems structure.
type GetItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetItemsOK creates a GetItemsOK with default headers values
func NewGetItemsOK() *GetItemsOK {
	return &GetItemsOK{}
}

/*
GetItemsOK describes a response with status code 200, with default header values.

OK
*/
type GetItemsOK struct {
	Payload *models.GetItemsResponse
}

// IsSuccess returns true when this get items o k response has a 2xx status code
func (o *GetItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get items o k response has a 3xx status code
func (o *GetItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get items o k response has a 4xx status code
func (o *GetItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get items o k response has a 5xx status code
func (o *GetItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get items o k response a status code equal to that given
func (o *GetItemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetItemsOK) Error() string {
	return fmt.Sprintf("[POST /items/get][%d] getItemsOK  %+v", 200, o.Payload)
}

func (o *GetItemsOK) String() string {
	return fmt.Sprintf("[POST /items/get][%d] getItemsOK  %+v", 200, o.Payload)
}

func (o *GetItemsOK) GetPayload() *models.GetItemsResponse {
	return o.Payload
}

func (o *GetItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.GetItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemsUnauthorized creates a GetItemsUnauthorized with default headers values
func NewGetItemsUnauthorized() *GetItemsUnauthorized {
	return &GetItemsUnauthorized{}
}

/*
GetItemsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetItemsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get items unauthorized response has a 2xx status code
func (o *GetItemsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get items unauthorized response has a 3xx status code
func (o *GetItemsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get items unauthorized response has a 4xx status code
func (o *GetItemsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get items unauthorized response has a 5xx status code
func (o *GetItemsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get items unauthorized response a status code equal to that given
func (o *GetItemsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetItemsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /items/get][%d] getItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetItemsUnauthorized) String() string {
	return fmt.Sprintf("[POST /items/get][%d] getItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetItemsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetItemsInternalServerError creates a GetItemsInternalServerError with default headers values
func NewGetItemsInternalServerError() *GetItemsInternalServerError {
	return &GetItemsInternalServerError{}
}

/*
GetItemsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetItemsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get items internal server error response has a 2xx status code
func (o *GetItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get items internal server error response has a 3xx status code
func (o *GetItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get items internal server error response has a 4xx status code
func (o *GetItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get items internal server error response has a 5xx status code
func (o *GetItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get items internal server error response a status code equal to that given
func (o *GetItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetItemsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /items/get][%d] getItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetItemsInternalServerError) String() string {
	return fmt.Sprintf("[POST /items/get][%d] getItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetItemsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
