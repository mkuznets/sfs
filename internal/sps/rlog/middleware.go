package rlog

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
	"net/http"
)

type contextKey int

const (
	ctxRequestIdKey = 2
)

// RequestIDHeader is the name of the HTTP Header which contains the request id.
// Exported so that it can be changed by developers
var RequestIDHeader = "X-Request-Id"

func GetReqID(r *http.Request) string {
	if reqID, ok := r.Context().Value(ctxRequestIdKey).(string); ok {
		return reqID
	}
	return ""
}

func RequestID(next http.Handler) http.Handler {
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

func Logger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if reqId := GetReqID(r); reqId != "" {
				logger := log.With().Str("req_id", reqId).Logger()
				ctx := logger.WithContext(r.Context())
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}
