package routes

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type GamesRouteHandler struct{}

func (GamesRouteHandler) Register(app fiber.Router) {
	gamesApi := app.Group("/games")

	gamesApi.Post("", createGame)
	gamesApi.Post("/:id/join", joinGame)
	gamesApi.Post("/:id/leave", leaveGame)
	gamesApi.Post("/:id/start", startGame)
	gamesApi.Post("/:id/turn", nextTurnGame)

	log.Println("Registered games api group.")
}

func createGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func joinGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func leaveGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func startGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func nextTurnGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
