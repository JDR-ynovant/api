package engine

import (
	"errors"
	"fmt"
	"github.com/JDR-ynovant/api/internal"
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/JDR-ynovant/api/internal/service/webpush"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PlayTurn(turn *models.Turn, game *models.Game) error {
	notificationStrings := internal.GetStrings()
	ur := repository.NewUserRepository()
	currentUser, _ := ur.FindOneById(game.Playing.Hex())

	if turn.Player != game.Playing {
		return errors.New("play turn: given turn character does not match with playing character")
	}

	game.Turns = append(game.Turns, models.Turn{
		Id:         primitive.NewObjectID(),
		X:          turn.X,
		Y:          turn.Y,
		Actions:    make([]models.Action, 0),
		Player:     turn.Player,
		TurnNumber: turn.TurnNumber,
	})

	for _, action := range turn.Actions {
		actionHandler := GetActionHandler(game, &action, turn)

		if isLegit, err := actionHandler.IsLegit(); !isLegit {
			return fmt.Errorf("play turn: action is not legal : %s", err.Error())
		}

		actionHandler.Handle()
	}

	// fetch dead player this turn for notification
	deadPlayers := getDeadPlayersThisTurn(game)
	_ = webpush.SendNotificationToPlayers(deadPlayers, fmt.Sprintf(notificationStrings.NotificationPlayerIsDead, currentUser.Name))

	// set new playing player
	game.Playing = calculateNewPlaying(game)

	// notify new playing user
	_ = webpush.SendNotificationToPlayer(*currentUser, notificationStrings.NotificationPlayerTurn)

	return nil
}

func getDeadPlayersThisTurn(game *models.Game) []primitive.ObjectID {
	config := internal.GetConfig()
	deadPlayers := make([]primitive.ObjectID, 0)

	for _, character := range game.Players {
		if character.BloodSugar == -1 {
			character.BloodSugar = config.RuleBloodSugarCap
			deadPlayers = append(deadPlayers, character.Id)
		}
	}

	return deadPlayers
}

func calculateNewPlaying(game *models.Game) primitive.ObjectID {
	// @todo
	// maybe a fifo loop ?
	//
	// or calculate based on last turns
	// array of excluded with current + dead players then rewind turns based on remaining players
	// and a map with pointer to player ID and last turn number played, once full get lowest turn number
	return primitive.ObjectID{}
}
