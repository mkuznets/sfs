package yerr

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	Status() int
	WithCause(error) Error
	Unwrap() error
}

type errorImpl struct {
	err     error
	status  int
	message string
}

func (e *errorImpl) Unwrap() error {
	return e.err
}

func (e *errorImpl) WithCause(err error) Error {
	e.err = err
	return e
}

func (e *errorImpl) Status() int {
	return e.status
}

func (e *errorImpl) Error() string {
	return e.message
}

func newErrorf(err error, status int, message string, a ...interface{}) Error {
	return &errorImpl{err, status, fmt.Sprintf(message, a...)}
}

func NotFound(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusNotFound, message, a...)
}

func Invalid(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusBadRequest, message, a...)
}

func Unauthorised(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusUnauthorized, message, a...)
}

func Internal(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusInternalServerError, message, a...)
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
