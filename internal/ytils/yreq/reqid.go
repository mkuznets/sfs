package yreq

import (
	"context"
	"github.com/segmentio/ksuid"
	"net/http"
)

type contextKey int

const (
	ctxRequestIdKey = contextKey(0x5245)
)

// RequestIDHeader is the name of the HTTP Header which contains the request id.
// Exported so that it can be changed by developers
var RequestIDHeader = "X-Request-Id"

func Id(r *http.Request) string {
	if reqID, ok := r.Context().Value(ctxRequestIdKey).(string); ok {
		return reqID
	}
	return ""
}

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = ksuid.New().String()
		}
		ctx = context.WithValue(ctx, ctxRequestIdKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
