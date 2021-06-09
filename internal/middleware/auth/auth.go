package auth

import (
	"fmt"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Get id from request, else we generate one
		user := c.Request().Header.Peek(cfg.Header)

		// Add the request ID to locals
		c.Locals(cfg.ContextKey, user)

		uString := fmt.Sprintf("%s", user)
		if uString != "" {
			ur := repository.NewUserRepository()
			userObject, err := ur.FindOneById(uString)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			c.Locals(cfg.ObjectKey, userObject)
		}

		// Continue stack
		return c.Next()
	}
}
