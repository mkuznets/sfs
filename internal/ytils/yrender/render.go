package yrender

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
	"net/http"
	"strings"
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

func renderJSONError(w http.ResponseWriter, r *http.Request, err error) {
	switch v := err.(type) {
	case yerr.Error:
		if v.Status() >= 500 {
			log.Ctx(r.Context()).Error().Stack().Err(v).Msg("server error")
		}

		var (
			msgs []string
			e    error
		)
		msgs = append(msgs, v.Error())

		if v.Status() < 500 {
			e = v
			for e != nil {
				e = errors.Unwrap(e)
				if e != nil {
					msgs = append(msgs, e.Error())
				}
			}
		}
		renderJSON(w, v.Status(), yerr.Response{Error: http.StatusText(v.Status()), Message: strings.Join(msgs, ": ")})
	default:
		renderJSONError(w, r, yerr.New("Internal error").Err(err))
	}
}
