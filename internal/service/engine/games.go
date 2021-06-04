package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"time"
)

const BASE_OBJECT_COUNT = 1.6

func GenerateGame(name string, owner string, playerCount int) (*models.Game, error) {
	ownerPrimitive, _ := primitive.ObjectIDFromHex(owner)
	grid, err := GenerateGrid(DEFAULT_GRID_WIDTH, DEFAULT_GRID_HEIGHT)

	if err != nil {
		return nil, err
	}

	objectCount := int(math.Round(BASE_OBJECT_COUNT * float64(playerCount)))
	objects := GenerateObjects(grid, objectCount)

	game := models.Game{
		Name:           name,
		Players:        make([]models.Character, 0),
		PlayerCount:    playerCount,
		Playing:        ownerPrimitive,
		Owner:          ownerPrimitive,
		Status:         models.GAME_STATUS_CREATED,
		Grid:           *grid,
		Objects:        objects,
		ExpiryDate:     time.Now().AddDate(0, 1, 0),
		Turns:          make([]models.Turn, 0),
		TurnNumber:     0,
		TurnExpiryDate: time.Now().AddDate(0, 0, 7),
	}

	return &game, nil
}
