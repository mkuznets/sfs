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

// GetChannelsIDReader is a Reader for the GetChannelsID structure.
type GetChannelsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChannelsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChannelsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetChannelsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetChannelsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChannelsIDOK creates a GetChannelsIDOK with default headers values
func NewGetChannelsIDOK() *GetChannelsIDOK {
	return &GetChannelsIDOK{}
}

/*
GetChannelsIDOK describes a response with status code 200, with default header values.

OK
*/
type GetChannelsIDOK struct {
	payload *models.ChannelResponse
}

// IsSuccess returns true when this get channels Id o k response has a 2xx status code
func (o *GetChannelsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get channels Id o k response has a 3xx status code
func (o *GetChannelsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get channels Id o k response has a 4xx status code
func (o *GetChannelsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get channels Id o k response has a 5xx status code
func (o *GetChannelsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get channels Id o k response a status code equal to that given
func (o *GetChannelsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetChannelsIDOK) Error() string {
	return fmt.Sprintf("[GET /channels/{id}][%d] getChannelsIdOK  %+v", 200, o.payload)
}

func (o *GetChannelsIDOK) String() string {
	return fmt.Sprintf("[GET /channels/{id}][%d] getChannelsIdOK  %+v", 200, o.payload)
}

func (o *GetChannelsIDOK) GetPayload() *models.ChannelResponse {
	return o.payload
}

func (o *GetChannelsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.payload = new(models.ChannelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChannelsIDNotFound creates a GetChannelsIDNotFound with default headers values
func NewGetChannelsIDNotFound() *GetChannelsIDNotFound {
	return &GetChannelsIDNotFound{}
}

/*
GetChannelsIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetChannelsIDNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get channels Id not found response has a 2xx status code
func (o *GetChannelsIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get channels Id not found response has a 3xx status code
func (o *GetChannelsIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get channels Id not found response has a 4xx status code
func (o *GetChannelsIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get channels Id not found response has a 5xx status code
func (o *GetChannelsIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get channels Id not found response a status code equal to that given
func (o *GetChannelsIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetChannelsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /channels/{id}][%d] getChannelsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetChannelsIDNotFound) String() string {
	return fmt.Sprintf("[GET /channels/{id}][%d] getChannelsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetChannelsIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetChannelsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChannelsIDInternalServerError creates a GetChannelsIDInternalServerError with default headers values
func NewGetChannelsIDInternalServerError() *GetChannelsIDInternalServerError {
	return &GetChannelsIDInternalServerError{}
}

/*
GetChannelsIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetChannelsIDInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get channels Id internal server error response has a 2xx status code
func (o *GetChannelsIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get channels Id internal server error response has a 3xx status code
func (o *GetChannelsIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get channels Id internal server error response has a 4xx status code
func (o *GetChannelsIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get channels Id internal server error response has a 5xx status code
func (o *GetChannelsIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get channels Id internal server error response a status code equal to that given
func (o *GetChannelsIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetChannelsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /channels/{id}][%d] getChannelsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChannelsIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /channels/{id}][%d] getChannelsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChannelsIDInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetChannelsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
