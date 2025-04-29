package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func serverTiming(c *fiber.Ctx) error {
	// Start the timer
	start := time.Now()

	// Process the request
	err := c.Next()

	// Stop the timer
	duration := time.Since(start)

	// Set the Server-Timing header
	c.Set("Server-Timing", strconv.FormatInt(duration.Milliseconds(), 10))

	return err
}

func ServerTimingMiddleware() fiber.Handler {
	return serverTiming
}
