package middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func IdentifierCookieMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("session_id")
			if err != nil || cookie.Value == "" {
				cookie := &http.Cookie{
					Name:     "session_id",
					Value:    uuid.NewString(),
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
					MaxAge:   int(time.Hour.Seconds()) * 24 * 30, // 30 days
				}
				c.SetCookie(cookie)
				c.Set("session_id", cookie)
			}
			return next(c)
		}
	}
}
