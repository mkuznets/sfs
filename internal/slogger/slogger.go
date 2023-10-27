package slogger

import (
	"context"
	"os"

	"github.com/mattn/go-isatty"
	"log/slog"
)

type ctxKey struct{}

func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
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
	slogger := slog.New(shandler)
	slog.SetDefault(slogger)
}
