package middleware

import (
	response "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/echo"
	"github.com/labstack/echo/v4"
)

type ValidatedHandler[T any] func(c echo.Context, req *T) error

func ValidateRequest[T any](handler ValidatedHandler[T]) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req T

		if err := c.Bind(&req); err != nil {
			return response.BadRequest(c, err)
		}

		if err := c.Validate(req); err != nil {
			return response.BadRequest(c, err)
		}

		return handler(c, &req)
	}
}
