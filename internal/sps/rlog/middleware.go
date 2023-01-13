package rlog

import (
	"context"
	"crypto/rand"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/http"
	"time"
)

type contextKey int

const (
	ctxLoggerKey    contextKey = 1
	ctxRequestIdKey            = 2
)

// RequestIDHeader is the name of the HTTP Header which contains the request id.
// Exported so that it can be changed by developers
var RequestIDHeader = "X-Request-Id"

func L(ctx context.Context) *zerolog.Logger {
	logger := ctx.Value(ctxLoggerKey)
	if logger == nil {
		return &log.Logger
	}
	if l, ok := logger.(*zerolog.Logger); ok {
		return l
	}
	return &log.Logger
}

func GetReqID(r *http.Request) string {
	if reqID, ok := r.Context().Value(ctxRequestIdKey).(string); ok {
		return reqID
	}
	return ""
}

func randBase62() string {
	var buf [12]byte
	_, _ = rand.Read(buf[:])
	var i big.Int
	i.SetBytes(buf[:])
	return i.Text(62)
}

func nowBase62() string {
	var i big.Int
	i.SetInt64(time.Now().UnixNano())
	return i.Text(62)
}

func genId() string {
	return nowBase62() + randBase62()
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = genId()
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
				ctx := r.Context()
				ctx = context.WithValue(ctx, ctxLoggerKey, &logger)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}
