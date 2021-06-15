package routes

import (
	"fmt"
	"github.com/JDR-ynovant/api/internal/middleware/auth"
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/JDR-ynovant/api/internal/service/engine"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type GamesRouteHandler struct{}

func (GamesRouteHandler) Register(app fiber.Router) {
	gamesApi := app.Group("/games")

	gamesApi.Post("", auth.NewAuthRequiredHandler(), handleCreateGame)
	gamesApi.Get("/:id", handleGetGame)
	gamesApi.Post("/:id/join", auth.NewAuthRequiredHandler(), handleJoinGame)
	gamesApi.Post("/:id/leave", auth.NewAuthRequiredHandler(), handleLeaveGame)
	gamesApi.Post("/:id/start", auth.NewAuthRequiredHandler(), handleStartGame)
	gamesApi.Post("/:id/stop", auth.NewAuthRequiredHandler(), handleStopGame)
	gamesApi.Post("/:id/turn", auth.NewAuthRequiredHandler(), handleNextTurn)

	log.Println("Registered games api group.")
}

type CreateGameRequest struct {
	Name        string
	PlayerCount int
}

func ValidateCreateGameRequest(gameRequest CreateGameRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(gameRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// handleGetGame godoc
// @Summary Get a Game
// @Description Fetch a game by its ID
// @Tags games
// @Accept  json
// @Produce  json
// @Param id path string true "Game ID"
// @Success 200 {object} models.Game
// @Router /api/games/{id} [get]
func handleGetGame(c *fiber.Ctx) error {
	gr := repository.NewGameRepository()

	game, err := gr.FindOneById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(game)
}

// handleCreateGame godoc
// @Summary Create a new Game
// @Description Generate all objects to generate a new Game
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "Owner of the game"
// @Param game body CreateGameRequest true "New Game data"
// @Success 200 {object} models.Game
// @Router /api/games [post]
func handleCreateGame(c *fiber.Ctx) error {
	owner := fmt.Sprintf("%s", c.Locals(auth.ContextKey))
	createGameRequest := new(CreateGameRequest)
	if err := c.BodyParser(createGameRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErrors := ValidateCreateGameRequest(*createGameRequest)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": validationErrors,
		})
	}

	gr := repository.NewGameRepository()
	game, _ := engine.GenerateGame(owner, createGameRequest.Name, createGameRequest.PlayerCount)

	createdGame, err := gr.Create(game)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = gr.AttachUser(owner, createdGame.Id.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(createdGame)
}

// handleJoinGame godoc
// @Summary Join a game
// @Description The given player will join the game
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "User to join the game"
// @Param id path string true "Game ID"
// @Success 200
// @Router /api/games/{id}/join [post]
func handleJoinGame(c *fiber.Ctx) error {
	player := fmt.Sprintf("%s", c.Locals(auth.ContextKey))
	playerObject := c.Locals(auth.ObjectKey).(*models.User)

	gr := repository.NewGameRepository()
	game, err := gr.FindOneById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if game.Status != models.GAME_STATUS_CREATED {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "game has already started",
		})
	}

	playerId, _ := primitive.ObjectIDFromHex(player)
	if !game.HasPlayer(playerId) {
		character := engine.CreateCharacter(playerId)
		game.Players = append(game.Players, *character)

		err = gr.Update(c.Params("id"), *game)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		ur := repository.NewUserRepository()
		playerObject.Games = append(playerObject.Games, game.Id)
		err = ur.Update(playerId.Hex(), *playerObject)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

// handleLeaveGame godoc
// @Summary Leave a game
// @Description The given player will leave the game
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "User to leave the game"
// @Param id path string true "Game ID"
// @Success 200
// @Router /api/games/{id}/leave [post]
func handleLeaveGame(c *fiber.Ctx) error {
	player := fmt.Sprintf("%s", c.Locals(auth.ContextKey))
	playerObject := c.Locals(auth.ObjectKey).(*models.User)

	gr := repository.NewGameRepository()
	game, err := gr.FindOneById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if game.Status == models.GAME_STATUS_FINISHED {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "game has finished",
		})
	}

	if game.Owner == playerObject.Id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "owner cannot leave its game",
		})
	}

	playerId, _ := primitive.ObjectIDFromHex(player)
	if game.HasPlayer(playerId) {
		game.RemovePlayer(playerObject.Id)
		err = gr.Update(c.Params("id"), *game)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		ur := repository.NewUserRepository()
		playerObject.RemoveGame(game.Id)
		err = ur.Update(playerId.Hex(), *playerObject)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

// handleStartGame godoc
// @Summary Start a game
// @Description The game will start, expiry date will be set and owners turn will begin. Only Game owner can start a Game.
// @Tags games
// @Accept json
// @Produce json
// @Param X-User header string true "Owner of the game to be started"
// @Param id path string true "Game ID"
// @Success 200
// @Router /api/games/{id}/start [post]
func handleStartGame(c *fiber.Ctx) error {
	playerObject := c.Locals(auth.ObjectKey).(*models.User)

	gameId := c.Params("id")
	if gameId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "missing game ID",
		})
	}

	gr := repository.NewGameRepository()
	game, err := gr.FindOneById(gameId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Game ID not found",
		})
	}

	if game.Owner != playerObject.Id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "only owner can start game its",
		})
	}

	grr := repository.NewGridRepository()
	grid := engine.GenerateGrid(engine.DEFAULT_GRID_WIDTH, engine.DEFAULT_GRID_HEIGHT)
	err = grr.Create(grid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	objects := engine.GenerateObjects(grid, game.PlayerCount)
	game.Objects = objects

	game.Grid = grid.Id
	game.Status = models.GAME_STATUS_STARTED

	err = gr.Update(gameId, *game)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

type NewTurnRequest struct {
	Actions []models.Action
	X       int
	Y       int
	Player  primitive.ObjectID
}

func ValidateNewTurnRequest(turnRequest NewTurnRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(turnRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// handleNextTurn godoc
// @Summary Play a games turn
// @Description Play the given turn of the given game.
// @Tags games
// @Accept json
// @Produce json
// @Param turn body NewTurnRequest true "New turn data"
// @Param id path string true "Game ID"
// @Success 200
// @Router /api/games/{id}/turn [post]
func handleNextTurn(c *fiber.Ctx) error {
	//playerObject := c.Locals(auth.ObjectKey).(*models.User)
	newTurnRequest := new(NewTurnRequest)
	if err := c.BodyParser(newTurnRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErrors := ValidateNewTurnRequest(*newTurnRequest)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": validationErrors,
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

// handleStopGame godoc
// @Summary Terminate a Game
// @Description The game will finish and no more turn can be played. Only Game owner can stop a Game.
// @Tags games
// @Accept json
// @Produce json
// @Param id path string true "Game ID"
// @Success 200
// @Router /api/games/{id}/stop [post]
func handleStopGame(c *fiber.Ctx) error {
	playerObject := c.Locals(auth.ObjectKey).(*models.User)

	gameId := c.Params("id")
	if gameId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "missing game ID",
		})
	}

	gr := repository.NewGameRepository()
	game, err := gr.FindOneById(gameId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Game ID not found",
		})
	}

	if game.Owner != playerObject.Id {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "only owner can start game its",
		})
	}

	game.Status = models.GAME_STATUS_FINISHED
	err = gr.Update(gameId, *game)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
