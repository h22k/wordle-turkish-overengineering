package middleware

import (
	"net/http"

	metrics "github.com/h22k/wordle-turkish-overengineering/server/internal/infrastructure/metric"
	"github.com/labstack/echo/v4"
)

func MetricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			status := c.Response().Status
			path := c.Path()
			method := c.Request().Method

			metrics.HttpRequests.WithLabelValues(method, path, http.StatusText(status)).Inc()

			return err
		}
	}
}
