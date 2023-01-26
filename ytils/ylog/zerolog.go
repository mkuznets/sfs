package ylog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sfs/ytils/yreq"
	"net/http"
	"os"
	"time"
)

func Setup() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	})
	zerolog.DefaultContextLogger = &log.Logger
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = false
	zerolog.ErrorStackMarshaler = MarshalStack
}

func ContextLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.Logger
		if reqId := yreq.Id(r); reqId != "" {
			logger = logger.With().Str("req_id", reqId).Logger()
		}
		next.ServeHTTP(w, r.WithContext(logger.WithContext(r.Context())))
	})
}
