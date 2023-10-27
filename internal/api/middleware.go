package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/segmentio/ksuid"
	"log/slog"

	"mkuznets.com/go/sfs/internal/slogger"
)

// RequestIDHeader is the name of the HTTP Header which contains the request id.
const RequestIDHeader = "X-Request-Id"

type ctxRequestIdKey struct{}

func RequestId(r *http.Request) string {
	if reqID, ok := r.Context().Value(ctxRequestIdKey{}).(string); ok {
		return reqID
	}
	return ""
}

func AddRequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = "req_" + ksuid.New().String()
		}
		ctx = context.WithValue(ctx, ctxRequestIdKey{}, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AddContextLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := slog.Default()
		if reqId := RequestId(r); reqId != "" {
			logger = logger.With("req_id", reqId)
		}

		ctx := slogger.NewContext(r.Context(), logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		requestURI := r.RequestURI

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		start := time.Now()
		defer func() {
			ctx := r.Context()
			logger := slogger.FromContext(ctx)

			logger.LogAttrs(ctx, slog.LevelInfo, "API",
				slog.String("method", method),
				slog.String("path", requestURI),
				slog.Duration("duration", time.Since(start)),
				slog.Int("status", ww.Status()),
				slog.Int("size", ww.BytesWritten()),
			)
		}()

		next.ServeHTTP(ww, r)
	})
}
