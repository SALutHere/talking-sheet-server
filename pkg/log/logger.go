package log

import (
	"io"
	"log/slog"
)

func New(cfg Config, out io.Writer) *slog.Logger {
	var h slog.Handler
	switch cfg.Env {
	case envLocal:
		h = slog.NewTextHandler(out, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case envDev:
		h = slog.NewJSONHandler(out, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case envProd:
		h = slog.NewJSONHandler(out, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	default:
		h = slog.NewTextHandler(out, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(h)
}
