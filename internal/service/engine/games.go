package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"time"
)

const BASE_OBJECT_COUNT = 1.6

func GenerateGame(owner *models.User, name string, playerCount int) (*models.Game, error) {
	players := []models.Character{*CreateCharacter(owner.Id, owner.Name)}

	game := models.Game{
		Name:           name,
		Players:        players,
		PlayerCount:    playerCount,
		Playing:        owner.Id,
		Owner:          owner.Id,
		Status:         models.GAME_STATUS_CREATED,
		Objects:        make([]models.Object, 0),
		ExpiryDate:     time.Now().AddDate(0, 1, 0),
		Turns:          make([]models.Turn, 0),
		TurnNumber:     0,
		TurnExpiryDate: time.Now().AddDate(0, 0, 7),
	}

	return &game, nil
}
