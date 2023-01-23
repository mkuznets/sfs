package yerr

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sps/internal/ytils/yrender"
	"net/http"
)

type Error interface {
	error
	Status() int
	Message() string
	WithStatus(int) Error
	WithError(error) Error
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

func (e *errorImpl) WithStatus(s int) Error {
	e.status = s
	return e
}

func (e *errorImpl) WithError(err error) Error {
	e.err = err
	return e
}

func (e *errorImpl) Status() int {
	return e.status
}

func (e *errorImpl) Message() string {
	return e.message
}

func (e *errorImpl) Error() string {
	return e.message
}

func newErrorf(err error, status int, message string, a ...interface{}) Error {
	return &errorImpl{err, status, fmt.Sprintf(message, a...)}
}

func Newf(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusInternalServerError, message, a...)
}

func Wrapf(err error, message string, a ...interface{}) Error {
	return newErrorf(err, http.StatusInternalServerError, message, a...)
}

func NotFound(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusNotFound, message, a...)
}

func Validation(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusBadRequest, message, a...)
}

func Unauthorised(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusUnauthorized, message, a...)
}

func Forbidden(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusForbidden, message, a...)
}

func Internal(message string, a ...interface{}) Error {
	return newErrorf(nil, http.StatusInternalServerError, message, a...)
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func RenderJson(w http.ResponseWriter, r *http.Request, err error) {
	switch v := err.(type) {
	case Error:
		if v.Status() >= 500 {
			log.Ctx(r.Context()).Error().Err(v.Unwrap()).Msg("server error")
		}
		yrender.Json(w, r, v.Status(), Response{Error: http.StatusText(v.Status()), Message: v.Message()})
	default:
		RenderJson(w, r, Internal("Internal error").WithError(err))
	}
}
