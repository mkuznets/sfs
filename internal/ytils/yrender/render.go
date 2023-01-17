package yrender

import (
	"bytes"
	"encoding/json"
	"net/http"
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
