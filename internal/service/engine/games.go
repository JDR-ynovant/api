package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const BASE_OBJECT_COUNT = 1.6

func GenerateGame(owner string, name string, playerCount int) (*models.Game, error) {
	ownerPrimitive, _ := primitive.ObjectIDFromHex(owner)

	game := models.Game{
		Name:           name,
		Players:        make([]models.Character, 0),
		PlayerCount:    playerCount,
		Playing:        ownerPrimitive,
		Owner:          ownerPrimitive,
		Status:         models.GAME_STATUS_CREATED,
		Objects:        make([]models.Object, 0),
		ExpiryDate:     time.Now().AddDate(0, 1, 0),
		Turns:          make([]models.Turn, 0),
		TurnNumber:     0,
		TurnExpiryDate: time.Now().AddDate(0, 0, 7),
	}

	return &game, nil
}
