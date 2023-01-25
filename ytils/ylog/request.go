package ylog

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func RequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			event := log.Ctx(r.Context()).Debug().
				Str("method", r.Method).
				Str("path", r.RequestURI)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				event.
					Dur("duration", time.Since(t1)).
					Int("status", ww.Status()).
					Int("size", ww.BytesWritten()).Msg("API")
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
