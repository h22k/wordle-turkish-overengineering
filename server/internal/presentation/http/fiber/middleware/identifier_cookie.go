package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func identifierCookie(c *fiber.Ctx) error {
	// Check if the identifier cookie is present
	cookie := c.Cookies("session_id", "")
	if cookie == "" {
		c.Cookie(&fiber.Cookie{
			Name:  "session_id",
			Value: uuid.NewString(),
		})
	}
	// Continue to the next middleware or handler
	return c.Next()
}

func IdentifierCookieMiddleware() fiber.Handler {
	return identifierCookie
}
