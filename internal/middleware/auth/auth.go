package auth

import (
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

		// Continue stack
		return c.Next()
	}
}