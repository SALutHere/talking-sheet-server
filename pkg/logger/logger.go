package logger

import (
	"io"
	"log/slog"
	"strings"
)

func New(cfg Config, out io.Writer) *slog.Logger {
	if out == nil {
		panic("logger: nil output writer")
	}

	env := strings.ToLower(strings.TrimSpace(cfg.Env))
	lvl := strings.ToLower(strings.TrimSpace(cfg.Level))

	opts := handlerOptions(lvl)
	h := handler(env, out, opts)

	return slog.New(h)
}

func handler(env string, out io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch env {
	case EnvProd, EnvDev:
		return slog.NewJSONHandler(out, opts)
	case EnvLocal:
		return slog.NewTextHandler(out, opts)
	default:
		return slog.NewTextHandler(out, opts)
	}
}

func handlerOptions(level string) *slog.HandlerOptions {
	switch level {
	case LvlWarn:
		return &slog.HandlerOptions{Level: slog.LevelWarn}
	case LvlInfo:
		return &slog.HandlerOptions{Level: slog.LevelInfo}
	case LvlDebug:
		return &slog.HandlerOptions{Level: slog.LevelDebug}
	case LvlError:
		return &slog.HandlerOptions{Level: slog.LevelError}
	default:
		return &slog.HandlerOptions{Level: slog.LevelInfo}
	}
}
