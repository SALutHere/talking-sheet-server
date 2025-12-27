package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const HeaderRequestID = "X-Request-Id"
const ctxKeyRequestID = "request_id"

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rid := c.Request().Header.Get(HeaderRequestID)
			if rid == "" {
				rid = uuid.NewString()
			}

			c.Set(ctxKeyRequestID, rid)
			c.Response().Header().Set(HeaderRequestID, rid)

			return next(c)
		}
	}
}

func GetRequestID(c echo.Context) string {
	if v, ok := c.Get(ctxKeyRequestID).(string); ok {
		return v
	}

	return ""
}
