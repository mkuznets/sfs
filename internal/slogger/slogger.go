package slogger

import (
	"context"
	"os"

	"github.com/mattn/go-isatty"
	"log/slog"
)

var _ slog.Handler = (*attrsHandler)(nil)

type (
	loggerKey struct{}
	attrsKey  struct{}
)

type loggerAttrs struct {
	args []any
}

func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, loggerKey{}, logger)
	ctx = context.WithValue(ctx, attrsKey{}, &loggerAttrs{})
	return ctx
}

func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}

func With(ctx context.Context, args ...any) {
	la, ok := ctx.Value(attrsKey{}).(*loggerAttrs)
	if !ok {
		panic("slogger: context not initialized")
	}
	la.args = append(la.args, args...)
}

func WithError(ctx context.Context, err error) {
	With(ctx, slog.Any("err", err))
}

type attrsHandler struct {
	slog.Handler
}

func (h *attrsHandler) Handle(ctx context.Context, r slog.Record) error {
	if l, ok := ctx.Value(attrsKey{}).(*loggerAttrs); ok {
		r.Add(l.args...)
	}
	return h.Handler.Handle(ctx, r)
}

func Init() {
	outputFile := os.Stderr

	hopts := &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}

	var shandler slog.Handler
	if fd := outputFile.Fd(); isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd) {
		shandler = slog.NewTextHandler(outputFile, hopts)
	} else {
		shandler = slog.NewJSONHandler(outputFile, hopts)
	}
	shandler = &attrsHandler{Handler: shandler}

	slogger := slog.New(shandler)
	slog.SetDefault(slogger)
}
