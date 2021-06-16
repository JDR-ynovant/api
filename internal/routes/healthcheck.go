package routes

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type HealthcheckRouteHandler struct{}

func (HealthcheckRouteHandler) Register(app fiber.Router) {
	app.Get("/healthcheck", handleHealthcheck)

	log.Println("Registered healthcheck api group.")
}

// handleSubscribe godoc
// @Summary Healthcheck route to ping API
// @Description Healthcheck route to ping API, will respond OK
// @Tags healthcheck
// @Accept  json
// @Produce  json
// @Success 200
// @Router /healthcheck [get]
func handleHealthcheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
