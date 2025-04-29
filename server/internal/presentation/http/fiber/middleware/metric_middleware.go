package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	metrics "github.com/h22k/wordle-turkish-overengineering/server/internal/metric"
)

func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		status := c.Response().StatusCode()
		path := c.Route().Path
		method := c.Method()

		metrics.HttpRequests.WithLabelValues(method, path, http.StatusText(status)).Inc()

		return err
	}
}
