package auth

import (
	"github.com/gofiber/fiber/v2"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Header is the header key where to get/set the unique request ID
	//
	// Optional. Default: "X-User"
	Header string

	// ContextKey defines the key used when storing the request ID in
	// the locals for a specific request.
	//
	// Optional. Default: authenticatedUser
	ContextKey string

	// ObjectKey defines the key used when storing the request ID in
	// the locals for a specific request.
	//
	// Optional. Default: authenticatedUser
	ObjectKey string
}

const ContextKey = "authenticatedUser"
const ObjectKey = "authenticatedUserObject"
const Header = "X-User"

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:       nil,
	Header:     Header,
	ContextKey: ContextKey,
	ObjectKey:  ObjectKey,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Header == "" {
		cfg.Header = ConfigDefault.Header
	}

	if cfg.ContextKey == "" {
		cfg.ContextKey = ConfigDefault.ContextKey
	}

	if cfg.ObjectKey == "" {
		cfg.ObjectKey = ConfigDefault.ObjectKey
	}

	return cfg
}
