// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mkuznets.com/go/sps/api/models"
)

// UploadFilesReader is a Reader for the UploadFiles structure.
type UploadFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadFilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadFilesOK creates a UploadFilesOK with default headers values
func NewUploadFilesOK() *UploadFilesOK {
	return &UploadFilesOK{}
}

/*
UploadFilesOK describes a response with status code 200, with default header values.

OK
*/
type UploadFilesOK struct {
	Payload *models.UploadFilesResponse
}

// IsSuccess returns true when this upload files o k response has a 2xx status code
func (o *UploadFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload files o k response has a 3xx status code
func (o *UploadFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload files o k response has a 4xx status code
func (o *UploadFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload files o k response has a 5xx status code
func (o *UploadFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload files o k response a status code equal to that given
func (o *UploadFilesOK) IsCode(code int) bool {
	return code == 200
}

func (o *UploadFilesOK) Error() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesOK  %+v", 200, o.Payload)
}

func (o *UploadFilesOK) String() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesOK  %+v", 200, o.Payload)
}

func (o *UploadFilesOK) GetPayload() *models.UploadFilesResponse {
	return o.Payload
}

func (o *UploadFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFilesBadRequest creates a UploadFilesBadRequest with default headers values
func NewUploadFilesBadRequest() *UploadFilesBadRequest {
	return &UploadFilesBadRequest{}
}

/*
UploadFilesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UploadFilesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this upload files bad request response has a 2xx status code
func (o *UploadFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload files bad request response has a 3xx status code
func (o *UploadFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload files bad request response has a 4xx status code
func (o *UploadFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload files bad request response has a 5xx status code
func (o *UploadFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload files bad request response a status code equal to that given
func (o *UploadFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UploadFilesBadRequest) Error() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesBadRequest  %+v", 400, o.Payload)
}

func (o *UploadFilesBadRequest) String() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesBadRequest  %+v", 400, o.Payload)
}

func (o *UploadFilesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFilesUnauthorized creates a UploadFilesUnauthorized with default headers values
func NewUploadFilesUnauthorized() *UploadFilesUnauthorized {
	return &UploadFilesUnauthorized{}
}

/*
UploadFilesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UploadFilesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this upload files unauthorized response has a 2xx status code
func (o *UploadFilesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload files unauthorized response has a 3xx status code
func (o *UploadFilesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload files unauthorized response has a 4xx status code
func (o *UploadFilesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload files unauthorized response has a 5xx status code
func (o *UploadFilesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upload files unauthorized response a status code equal to that given
func (o *UploadFilesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UploadFilesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesUnauthorized  %+v", 401, o.Payload)
}

func (o *UploadFilesUnauthorized) String() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesUnauthorized  %+v", 401, o.Payload)
}

func (o *UploadFilesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadFilesInternalServerError creates a UploadFilesInternalServerError with default headers values
func NewUploadFilesInternalServerError() *UploadFilesInternalServerError {
	return &UploadFilesInternalServerError{}
}

/*
UploadFilesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UploadFilesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this upload files internal server error response has a 2xx status code
func (o *UploadFilesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload files internal server error response has a 3xx status code
func (o *UploadFilesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload files internal server error response has a 4xx status code
func (o *UploadFilesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload files internal server error response has a 5xx status code
func (o *UploadFilesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload files internal server error response a status code equal to that given
func (o *UploadFilesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UploadFilesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadFilesInternalServerError) String() string {
	return fmt.Sprintf("[POST /files/upload][%d] uploadFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadFilesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UploadFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}