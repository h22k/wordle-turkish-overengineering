package response

import (
	"github.com/gofiber/fiber/v2"
)

func StreamResponse(c *fiber.Ctx, data interface{}) {

}

func Success(c *fiber.Ctx, data interface{}) error {
	return success(c, fiber.StatusOK, data)
}

func Created(c *fiber.Ctx, data interface{}) error {
	return success(c, fiber.StatusCreated, data)
}

func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

func BadRequest(c *fiber.Ctx, error error) error {
	return err(c, fiber.StatusBadRequest, error)
}

func success(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"data": data,
	})
}

func err(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}
