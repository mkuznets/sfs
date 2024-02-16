package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/segmentio/ksuid"
	"log/slog"

	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/slogger"
)

// RequestIDHeader is the name of the HTTP Header which contains the request id.
const RequestIDHeader = "X-Request-Id"

type ctxRequestIdKey struct{}

const maxHeaderLength = 512

var sensitiveHeaders = map[string]struct{}{
	"authorization": {},
	"cookie":        {},
	"x-csrf-token":  {},
}

func sanitizeHeader(key string, values []string) any {
	if _, ok := sensitiveHeaders[key]; ok {
		values = []string{"[redacted]"}
	}

	for i, v := range values {
		if len(v) > maxHeaderLength {
			values[i] = v[:maxHeaderLength] + "..."
		}
	}

	if len(values) == 1 {
		return values[0]
	}

	return values
}

func RequestID(r *http.Request) string {
	if reqID, ok := r.Context().Value(ctxRequestIdKey{}).(string); ok {
		return reqID
	}
	return ""
}

func AddRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = "req_" + ksuid.New().String()
		}
		slogger.With(ctx, slog.String("req_id", requestID))
		ctx = context.WithValue(ctx, ctxRequestIdKey{}, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AddContextLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := slogger.NewContext(r.Context(), slog.Default())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.UserFromContext(r.Context())
		slogger.With(r.Context(), slog.String("user_id", user.ID()))
		next.ServeHTTP(w, r)
	})
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logAttrs := []slog.Attr{
			slog.String("method", r.Method),
			slog.String("request_uri", r.RequestURI),
			slog.String("proto", r.Proto),
			slog.String("host", r.Host),
			slog.Int64("content_length", r.ContentLength),
		}

		for key, values := range r.Header {
			key = strings.ToLower(key)
			attrKey := fmt.Sprintf("headers.%s", key)
			safeValue := sanitizeHeader(key, values)
			logAttrs = append(logAttrs, slog.Any(attrKey, safeValue))
		}

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		start := time.Now()
		next.ServeHTTP(ww, r)
		duration := time.Since(start)

		logAttrs = append(logAttrs,
			slog.String("remote_addr", r.RemoteAddr),
			slog.Duration("duration", duration),
			slog.Int("response_status", ww.Status()),
			slog.Int("response_size", ww.BytesWritten()),
		)

		level := slog.LevelInfo
		if ww.Status() >= 500 {
			level = slog.LevelError
		}

		ctx := r.Context()
		lgr := slogger.FromContext(ctx)
		lgr.LogAttrs(ctx, level, "API", logAttrs...)
	})
}
