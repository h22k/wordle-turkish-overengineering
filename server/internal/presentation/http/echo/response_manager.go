package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data interface{}) error {
	return success(c, http.StatusOK, data)
}

func Created(c echo.Context, data interface{}) error {
	return success(c, http.StatusCreated, data)
}

func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func BadRequest(c echo.Context, err error) error {
	return errResponse(c, http.StatusBadRequest, err)
}

func success(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, map[string]interface{}{
		"data": data,
	})
}

func errResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, map[string]interface{}{
		"error": err.Error(),
	})
}
