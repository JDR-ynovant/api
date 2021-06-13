package auth

import (
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/gofiber/fiber/v2"
)

var moduleConfiguration Config

// NewAuthHeaderHandler creates a new middleware handler
func NewAuthHeaderHandler(config ...Config) fiber.Handler {
	// Set default config
	moduleConfiguration = configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if moduleConfiguration.Next != nil && moduleConfiguration.Next(c) {
			return c.Next()
		}

		// Get id from request, else we generate one
		user := c.Request().Header.Peek(moduleConfiguration.Header)

		// Add the request ID to locals
		c.Locals(moduleConfiguration.ContextKey, user)

		uString := string(user)
		if uString != "" {
			ur := repository.NewUserRepository()
			userObject, err := ur.FindOneById(uString)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": err.Error(),
				})
			}

			c.Locals(moduleConfiguration.ObjectKey, userObject)
		}

		// Continue stack
		return c.Next()
	}
}

// NewAuthRequiredHandler creates a new middleware handler
func NewAuthRequiredHandler() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if moduleConfiguration.Next != nil && moduleConfiguration.Next(c) {
			return c.Next()
		}

		playerId := string(c.Locals(moduleConfiguration.ContextKey).([]byte))
		if playerId == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Continue stack
		return c.Next()
	}
}
