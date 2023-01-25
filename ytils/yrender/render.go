package yrender

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sfs/ytils/yerr"
	"net/http"
)

type Response interface {
	Status(s int) Response
	JSON()
	XML()
}

type response struct {
	w      http.ResponseWriter
	r      *http.Request
	status int
	v      interface{}
}

func New(w http.ResponseWriter, r *http.Request, v interface{}) Response {
	return &response{
		w:      w,
		r:      r,
		v:      v,
		status: http.StatusOK,
	}
}

func (r *response) Status(status int) Response {
	r.status = status
	return r
}

func (r *response) JSON() {
	switch obj := r.v.(type) {
	case error:
		renderJSONError(r.w, r.r, obj)
	default:
		renderJSON(r.w, r.status, r.v)
	}
}

func (r *response) XML() {
	r.w.Header().Set("Content-Type", "text/xml; charset=utf-8")

	switch obj := r.v.(type) {
	case string:
		r.w.WriteHeader(r.status)
		_, _ = r.w.Write([]byte(obj))
	default:
		http.Error(r.w, "Unsupported type", http.StatusInternalServerError)
	}
}

func renderJSON(w http.ResponseWriter, status int, v interface{}) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}

func reportError(ctx context.Context, err error, stack bool) {
	event := log.Ctx(ctx).Error()
	if stack {
		event = event.Stack()
	}
	event.Err(err).Send()
}

func renderJSONError(w http.ResponseWriter, r *http.Request, err error) {
	switch v := err.(type) {
	case yerr.Error:
		message := v.Error()
		stack := false
		if v.Status() >= 500 {
			message = "Internal Server Error"
			stack = true
		}
		reportError(r.Context(), err, stack)
		renderJSON(w, v.Status(), ErrorResponse{
			Error:   http.StatusText(v.Status()),
			Message: message,
		})
	default:
		reportError(r.Context(), err, false)
		renderJSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "Internal Server Error",
		})
	}
}
