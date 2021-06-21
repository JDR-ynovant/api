package routes

import "github.com/gofiber/fiber/v2"

type RouteHandler interface {
	Register(app fiber.Router)
}

func jsonError(c *fiber.Ctx, errorCode int, errorMessage interface{}) error {
	return c.Status(errorCode).JSON(fiber.Map{
		"message": errorMessage,
	})
}
