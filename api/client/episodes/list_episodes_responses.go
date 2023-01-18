// Code generated by go-swagger; DO NOT EDIT.

package episodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mkuznets.com/go/sps/api/models"
)

// ListEpisodesReader is a Reader for the ListEpisodes structure.
type ListEpisodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEpisodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEpisodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEpisodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListEpisodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListEpisodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListEpisodesOK creates a ListEpisodesOK with default headers values
func NewListEpisodesOK() *ListEpisodesOK {
	return &ListEpisodesOK{}
}

/*
ListEpisodesOK describes a response with status code 200, with default header values.

OK
*/
type ListEpisodesOK struct {
	Payload *models.IDResponse
}

// IsSuccess returns true when this list episodes o k response has a 2xx status code
func (o *ListEpisodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list episodes o k response has a 3xx status code
func (o *ListEpisodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list episodes o k response has a 4xx status code
func (o *ListEpisodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list episodes o k response has a 5xx status code
func (o *ListEpisodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list episodes o k response a status code equal to that given
func (o *ListEpisodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListEpisodesOK) Error() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesOK  %+v", 200, o.Payload)
}

func (o *ListEpisodesOK) String() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesOK  %+v", 200, o.Payload)
}

func (o *ListEpisodesOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *ListEpisodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEpisodesUnauthorized creates a ListEpisodesUnauthorized with default headers values
func NewListEpisodesUnauthorized() *ListEpisodesUnauthorized {
	return &ListEpisodesUnauthorized{}
}

/*
ListEpisodesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListEpisodesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list episodes unauthorized response has a 2xx status code
func (o *ListEpisodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list episodes unauthorized response has a 3xx status code
func (o *ListEpisodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list episodes unauthorized response has a 4xx status code
func (o *ListEpisodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list episodes unauthorized response has a 5xx status code
func (o *ListEpisodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list episodes unauthorized response a status code equal to that given
func (o *ListEpisodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListEpisodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListEpisodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListEpisodesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEpisodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEpisodesNotFound creates a ListEpisodesNotFound with default headers values
func NewListEpisodesNotFound() *ListEpisodesNotFound {
	return &ListEpisodesNotFound{}
}

/*
ListEpisodesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListEpisodesNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list episodes not found response has a 2xx status code
func (o *ListEpisodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list episodes not found response has a 3xx status code
func (o *ListEpisodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list episodes not found response has a 4xx status code
func (o *ListEpisodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list episodes not found response has a 5xx status code
func (o *ListEpisodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list episodes not found response a status code equal to that given
func (o *ListEpisodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListEpisodesNotFound) Error() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesNotFound  %+v", 404, o.Payload)
}

func (o *ListEpisodesNotFound) String() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesNotFound  %+v", 404, o.Payload)
}

func (o *ListEpisodesNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEpisodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEpisodesInternalServerError creates a ListEpisodesInternalServerError with default headers values
func NewListEpisodesInternalServerError() *ListEpisodesInternalServerError {
	return &ListEpisodesInternalServerError{}
}

/*
ListEpisodesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ListEpisodesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this list episodes internal server error response has a 2xx status code
func (o *ListEpisodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list episodes internal server error response has a 3xx status code
func (o *ListEpisodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list episodes internal server error response has a 4xx status code
func (o *ListEpisodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list episodes internal server error response has a 5xx status code
func (o *ListEpisodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list episodes internal server error response a status code equal to that given
func (o *ListEpisodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListEpisodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListEpisodesInternalServerError) String() string {
	return fmt.Sprintf("[GET /channels/{id}/episodes][%d] listEpisodesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListEpisodesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEpisodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
