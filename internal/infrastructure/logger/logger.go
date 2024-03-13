package logger

import (
	"context"
	"log/slog"
	"os"
)

type logCtxKey struct{}

func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func AddToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logCtxKey{}, New())
}

func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(logCtxKey{}).(*slog.Logger); ok {
		return logger
	}
	return New()
}
