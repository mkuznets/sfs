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

	"mkuznets.com/go/sfs/internal/slogger"
)

// RequestIDHeader is the name of the HTTP Header which contains the request id.
const RequestIDHeader = "X-Request-Id"

type ctxRequestIdKey struct{}

var sensitiveHeaders = map[string]struct{}{
	"authorization": {},
	"cookie":        {},
	"x-csrf-token":  {},
}

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
		logAttrs := []slog.Attr{
			slog.String("method", r.Method),
			slog.String("request_uri", r.RequestURI),
			slog.String("proto", r.Proto),
			slog.String("host", r.Host),
			slog.String("remote_addr", r.RemoteAddr),
			slog.Int64("content_length", r.ContentLength),
		}

		for key, values := range r.Header {
			key = strings.ToLower(key)
			if _, ok := sensitiveHeaders[key]; ok {
				values = []string{"[redacted]"}
			}
			attrKey := fmt.Sprintf("headers.%s", key)
			var attrValue any
			if len(values) == 1 {
				attrValue = values[0]
			} else {
				attrValue = values
			}
			logAttrs = append(logAttrs, slog.Any(attrKey, attrValue))
		}

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		t1 := time.Now()
		defer func() {
			logAttrs = append(logAttrs,
				slog.Duration("duration", time.Since(t1)),
				slog.Int("response_status", ww.Status()),
				slog.Int("response_size", ww.BytesWritten()),
			)

			ctx := r.Context()
			lgr := slogger.FromContext(ctx)
			lgr.LogAttrs(ctx, slog.LevelInfo, "API", logAttrs...)
		}()

		next.ServeHTTP(ww, r)
	})
}
