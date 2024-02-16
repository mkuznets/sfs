package render

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

var statusCodeRegex = regexp.MustCompile(`^HTTP (\d{3}):\s*(.+)`)

type Response struct {
	writer  http.ResponseWriter
	request *http.Request

	status int
	value  interface{}
}

func New(w http.ResponseWriter, r *http.Request, v any) *Response {
	return &Response{
		writer:  w,
		request: r,
		value:   v,
		status:  http.StatusOK,
	}
}

func (r *Response) Status(status int) *Response {
	r.status = status
	return r
}

func (r *Response) JSON() {
	switch value := r.value.(type) {
	case error:
		fallbackStatus := http.StatusInternalServerError
		if r.status != http.StatusOK {
			fallbackStatus = r.status
		}
		code, msg := extractStatus(value, fallbackStatus)
		if code >= 500 {
			msg = "Internal Server Error"
		}

		renderJSON(r.writer, code, errorResponse{
			Error:   http.StatusText(code),
			Message: msg,
		})
	default:
		renderJSON(r.writer, r.status, r.value)
	}
}

func (r *Response) XML() {
	r.writer.Header().Set("Content-Type", "text/xml; charset=utf-8")

	switch obj := r.value.(type) {
	case string:
		r.writer.WriteHeader(r.status)
		_, _ = r.writer.Write([]byte(obj))
	default:
		http.Error(r.writer, "Unsupported type", http.StatusInternalServerError)
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

func extractStatus(err error, fallback int) (int, string) {
	msg := err.Error()
	if matches := statusCodeRegex.FindStringSubmatch(msg); matches != nil {
		statusCode, _ := strconv.Atoi(matches[1])
		msg = matches[2]
		return statusCode, msg
	}
	return fallback, msg
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
