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

// CreateEpisodeReader is a Reader for the CreateEpisode structure.
type CreateEpisodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEpisodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEpisodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEpisodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEpisodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEpisodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEpisodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEpisodeOK creates a CreateEpisodeOK with default headers values
func NewCreateEpisodeOK() *CreateEpisodeOK {
	return &CreateEpisodeOK{}
}

/*
CreateEpisodeOK describes a response with status code 200, with default header values.

OK
*/
type CreateEpisodeOK struct {
	Payload *models.IDResponse
}

// IsSuccess returns true when this create episode o k response has a 2xx status code
func (o *CreateEpisodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create episode o k response has a 3xx status code
func (o *CreateEpisodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create episode o k response has a 4xx status code
func (o *CreateEpisodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create episode o k response has a 5xx status code
func (o *CreateEpisodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create episode o k response a status code equal to that given
func (o *CreateEpisodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateEpisodeOK) Error() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeOK  %+v", 200, o.Payload)
}

func (o *CreateEpisodeOK) String() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeOK  %+v", 200, o.Payload)
}

func (o *CreateEpisodeOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *CreateEpisodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEpisodeBadRequest creates a CreateEpisodeBadRequest with default headers values
func NewCreateEpisodeBadRequest() *CreateEpisodeBadRequest {
	return &CreateEpisodeBadRequest{}
}

/*
CreateEpisodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateEpisodeBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create episode bad request response has a 2xx status code
func (o *CreateEpisodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create episode bad request response has a 3xx status code
func (o *CreateEpisodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create episode bad request response has a 4xx status code
func (o *CreateEpisodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create episode bad request response has a 5xx status code
func (o *CreateEpisodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create episode bad request response a status code equal to that given
func (o *CreateEpisodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateEpisodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEpisodeBadRequest) String() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEpisodeBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEpisodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEpisodeUnauthorized creates a CreateEpisodeUnauthorized with default headers values
func NewCreateEpisodeUnauthorized() *CreateEpisodeUnauthorized {
	return &CreateEpisodeUnauthorized{}
}

/*
CreateEpisodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateEpisodeUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create episode unauthorized response has a 2xx status code
func (o *CreateEpisodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create episode unauthorized response has a 3xx status code
func (o *CreateEpisodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create episode unauthorized response has a 4xx status code
func (o *CreateEpisodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create episode unauthorized response has a 5xx status code
func (o *CreateEpisodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create episode unauthorized response a status code equal to that given
func (o *CreateEpisodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateEpisodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateEpisodeUnauthorized) String() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateEpisodeUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEpisodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEpisodeNotFound creates a CreateEpisodeNotFound with default headers values
func NewCreateEpisodeNotFound() *CreateEpisodeNotFound {
	return &CreateEpisodeNotFound{}
}

/*
CreateEpisodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateEpisodeNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create episode not found response has a 2xx status code
func (o *CreateEpisodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create episode not found response has a 3xx status code
func (o *CreateEpisodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create episode not found response has a 4xx status code
func (o *CreateEpisodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create episode not found response has a 5xx status code
func (o *CreateEpisodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create episode not found response a status code equal to that given
func (o *CreateEpisodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateEpisodeNotFound) Error() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeNotFound  %+v", 404, o.Payload)
}

func (o *CreateEpisodeNotFound) String() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeNotFound  %+v", 404, o.Payload)
}

func (o *CreateEpisodeNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEpisodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEpisodeInternalServerError creates a CreateEpisodeInternalServerError with default headers values
func NewCreateEpisodeInternalServerError() *CreateEpisodeInternalServerError {
	return &CreateEpisodeInternalServerError{}
}

/*
CreateEpisodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateEpisodeInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create episode internal server error response has a 2xx status code
func (o *CreateEpisodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create episode internal server error response has a 3xx status code
func (o *CreateEpisodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create episode internal server error response has a 4xx status code
func (o *CreateEpisodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create episode internal server error response has a 5xx status code
func (o *CreateEpisodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create episode internal server error response a status code equal to that given
func (o *CreateEpisodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateEpisodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEpisodeInternalServerError) String() string {
	return fmt.Sprintf("[POST /channels/{id}/episodes][%d] createEpisodeInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEpisodeInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEpisodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
