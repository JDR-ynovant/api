package routes

import (
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type CreateGameRequest struct {
	Name        string
	PlayerCount int
}

// createGame godoc
// @Summary Create a new Game
// @Description Generate all objects to generate a new Game
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "Owner of the game"
// @Param game body CreateGameRequest true "New Game data"
// @Success 200 {object} models.Game
// @Router /api/games [post]
func createGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// joinGame godoc
// @Summary Join a game
// @Description The given player will join the game
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "User to join the game"
// @Param id path int true "Game ID"
// @Success 200
// @Router /api/games/{id}/join [post]
func joinGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// leaveGame godoc
// @Summary Join a game
// @Description The given player will leave the game
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "User to leave the game"
// @Param id path int true "Game ID"
// @Success 200
// @Router /api/games/{id}/leave [post]
func leaveGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// leaveGame godoc
// @Summary Start a game
// @Description The game will start, expiry date will be set and owners turn will begin.
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "Owner of the game to be started"
// @Param id path int true "Game ID"
// @Success 200
// @Router /api/games/{id}/start [post]
func startGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

type NewTurnRequest struct {
	Actions []models.Action
	X       int
	Y       int
	Player  primitive.ObjectID
}

// nextTurnGame godoc
// @Summary Play a games turn
// @Description Play the given turn of the given game.
// @Tags games
// @Accept json
// @Produce json
// @Param turn body NewTurnRequest true "New turn data"
// @Param id path int true "Game ID"
// @Success 200
// @Router /api/games/{id}/turn [post]
func nextTurnGame(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
