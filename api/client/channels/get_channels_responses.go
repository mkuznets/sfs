// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mkuznets.com/go/sps/api/models"
)

// GetChannelsReader is a Reader for the GetFeeds structure.
type GetChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetChannelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetChannelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChannelsOK creates a GetChannelsOK with default headers values
func NewGetChannelsOK() *GetChannelsOK {
	return &GetChannelsOK{}
}

/*
GetChannelsOK describes a response with status code 200, with default header values.

OK
*/
type GetChannelsOK struct {
	Payload *models.IDResponse
}

// IsSuccess returns true when this get channels o k response has a 2xx status code
func (o *GetChannelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get channels o k response has a 3xx status code
func (o *GetChannelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get channels o k response has a 4xx status code
func (o *GetChannelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get channels o k response has a 5xx status code
func (o *GetChannelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get channels o k response a status code equal to that given
func (o *GetChannelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetChannelsOK) Error() string {
	return fmt.Sprintf("[GET /channels][%d] getChannelsOK  %+v", 200, o.Payload)
}

func (o *GetChannelsOK) String() string {
	return fmt.Sprintf("[GET /channels][%d] getChannelsOK  %+v", 200, o.Payload)
}

func (o *GetChannelsOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *GetChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChannelsUnauthorized creates a GetChannelsUnauthorized with default headers values
func NewGetChannelsUnauthorized() *GetChannelsUnauthorized {
	return &GetChannelsUnauthorized{}
}

/*
GetChannelsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetChannelsUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get channels unauthorized response has a 2xx status code
func (o *GetChannelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get channels unauthorized response has a 3xx status code
func (o *GetChannelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get channels unauthorized response has a 4xx status code
func (o *GetChannelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get channels unauthorized response has a 5xx status code
func (o *GetChannelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get channels unauthorized response a status code equal to that given
func (o *GetChannelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetChannelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /channels][%d] getChannelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChannelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /channels][%d] getChannelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChannelsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetChannelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChannelsInternalServerError creates a GetChannelsInternalServerError with default headers values
func NewGetChannelsInternalServerError() *GetChannelsInternalServerError {
	return &GetChannelsInternalServerError{}
}

/*
GetChannelsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetChannelsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get channels internal server error response has a 2xx status code
func (o *GetChannelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get channels internal server error response has a 3xx status code
func (o *GetChannelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get channels internal server error response has a 4xx status code
func (o *GetChannelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get channels internal server error response has a 5xx status code
func (o *GetChannelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get channels internal server error response a status code equal to that given
func (o *GetChannelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetChannelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /channels][%d] getChannelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChannelsInternalServerError) String() string {
	return fmt.Sprintf("[GET /channels][%d] getChannelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChannelsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetChannelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}