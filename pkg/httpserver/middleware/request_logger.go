package middleware

import (
	"log/slog"
	"time"

	"github.com/SALutHere/talking-sheet-server/pkg/logger"
	"github.com/labstack/echo/v4"
)

func RequestLogger(log *slog.Logger) echo.MiddlewareFunc {
	if log == nil {
		log = slog.Default()
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			req := c.Request()
			res := c.Response()

			rid := GetRequestID(c)

			reqLog := log.With(
				slog.String("request_id", rid),
				slog.String("method", req.Method),
				slog.String("path", req.URL.Path),
				slog.String("route", c.Path()),
			)

			ctx := logger.With(req.Context(), reqLog)
			c.SetRequest(req.WithContext(ctx))

			reqLog.Info("request started")

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			reqLog.Info("request finished",
				slog.Int("status", res.Status),
				slog.Int64("bytes_out", res.Size),
				slog.Duration("latency", time.Since(start)),
			)

			return err
		}
	}
}
