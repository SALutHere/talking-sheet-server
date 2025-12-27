package middleware

import (
	"log/slog"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func Recover(log *slog.Logger) echo.MiddlewareFunc {
	if log == nil {
		log = slog.Default()
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					log.Error("panic recovered",
						slog.Any("panic", r),
						slog.String("request_id", GetRequestID(c)),
						slog.String("method", c.Request().Method),
						slog.String("path", c.Request().URL.Path),
						slog.String("route", c.Path()),
						slog.String("stack", string(debug.Stack())),
					)
					err = echo.NewHTTPError(500, "internal server error")
				}
			}()

			return next(c)
		}
	}
}
