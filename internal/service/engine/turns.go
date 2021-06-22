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

	// determine if game victory condition is reached
	if isVictoryConditionReached(game) {
		game.Status = models.GAME_STATUS_FINISHED
		_ = webpush.SendNotificationToPlayer(*currentUser, notificationStrings.NotificationPlayerWin)
		return nil
	}

	// set new playing player
	game.Playing = calculateNewPlaying(game)

	// notify new playing user
	_ = webpush.SendNotificationToPlayer(*currentUser, notificationStrings.NotificationPlayerTurn)

	return nil
}

func isVictoryConditionReached(game *models.Game) bool {
	// victory condition are : the playing user is the last alive

	return true
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
	config := internal.GetConfig()
	turnCharacterMap := make(map[primitive.ObjectID]int)
	excludedCharacter := make([]primitive.ObjectID, 0)
	excludedCharacter = append(excludedCharacter, game.Playing)

	// exclude dead players
	for _, character := range game.Players {
		if character.BloodSugar == config.RuleBloodSugarCap {
			excludedCharacter = append(excludedCharacter, character.Id)
		}
	}

	// init turnCharacterMap with all non excluded players
	for _, character := range game.Players {
		if isPlayerExcluded := inArray(excludedCharacter, character.Id); !isPlayerExcluded {
			turnCharacterMap[character.Id] = -1
		}
	}

	for i := len(game.Turns) - 1; i >= 0; i-- {
		//fmt.Printf("inspecting turn : %v for game : %v\n", i, game.Id)
		turn := game.Turns[i]
		if isPlayerExcluded := inArray(excludedCharacter, turn.Player); !isPlayerExcluded {
			turnCharacterMap[turn.Player] = turn.TurnNumber
		}
	}

	// iterate over map and return lowest turn number
	selectedNewPlaying := game.Playing
	lowestTurnPlayed := game.TurnNumber
	for id, lastTurnPlayed := range turnCharacterMap {
		if lastTurnPlayed < lowestTurnPlayed {
			selectedNewPlaying = id
			lowestTurnPlayed = lastTurnPlayed
		}
		//fmt.Printf("calculateNewPlaying: player %v : last turn %v; selected = %v\n", id, lastTurnPlayed, selectedNewPlaying)
	}

	return selectedNewPlaying
}

func inArray(array []primitive.ObjectID, needle primitive.ObjectID) bool {
	for _, id := range array {
		if id == needle {
			return true
		}
	}
	return false
}