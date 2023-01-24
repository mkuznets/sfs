// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mkuznets.com/go/sps/api/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserOK creates a CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {
	return &CreateUserOK{}
}

/*
CreateUserOK describes a response with status code 200, with default header values.

OK
*/
type CreateUserOK struct {
	Payload *models.IDResponse
}

// IsSuccess returns true when this create user o k response has a 2xx status code
func (o *CreateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user o k response has a 3xx status code
func (o *CreateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user o k response has a 4xx status code
func (o *CreateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user o k response has a 5xx status code
func (o *CreateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create user o k response a status code equal to that given
func (o *CreateUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateUserOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserOK  %+v", 200, o.Payload)
}

func (o *CreateUserOK) String() string {
	return fmt.Sprintf("[POST /users][%d] createUserOK  %+v", 200, o.Payload)
}

func (o *CreateUserOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *CreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/*
CreateUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateUserInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create user internal server error response has a 2xx status code
func (o *CreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user internal server error response has a 3xx status code
func (o *CreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user internal server error response has a 4xx status code
func (o *CreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user internal server error response has a 5xx status code
func (o *CreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user internal server error response a status code equal to that given
func (o *CreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /users][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}