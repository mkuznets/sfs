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

func Json(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}

func Rss(w http.ResponseWriter, r *http.Request, status int, v string) {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(v))
}

func JsonErr(w http.ResponseWriter, r *http.Request, err error) {
	switch v := err.(type) {
	case yerr.Error:
		if v.Status() >= 500 {
			log.Ctx(r.Context()).Error().Err(v.(error)).Msg("server error")
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
		Json(w, r, v.Status(), yerr.Response{Error: http.StatusText(v.Status()), Message: strings.Join(msgs, ": ")})
	default:
		JsonErr(w, r, yerr.Internal("Internal error").WithCause(err))
	}
}
