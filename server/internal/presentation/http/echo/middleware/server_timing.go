package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ServerTimingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			duration := time.Since(start)
			c.Response().Header().Set("Server-Timing", strconv.FormatInt(duration.Milliseconds(), 10))

			return err
		}
	}
}
