package engine

import (
	"fmt"
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCharacter(user primitive.ObjectID, name string) *models.Character {
	character := models.Character{
		Id:         primitive.NewObjectID(),
		Name:       name,
		User:       user,
		BloodSugar: 0,
		Inventory:  make([]models.Object, 0),
		PositionX:  0,
		PositionY:  0,
	}

	return &character
}

func GeneratePlayersPosition(grid models.Grid, game models.Game) []models.Character {
	characters := make([]models.Character, 0)
	playerMinimalRange := getPlayerMinimalRange(game.PlayerCount, grid.Height, grid.Width)

	for _, character := range game.Players {
		character.PositionX, character.PositionY = randomPositionOverRange(characters, playerMinimalRange, grid, game.Objects)
		characters = append(characters, character)
	}

	return characters
}

func randomPositionOverRange(characters []models.Character, minimalRange int, grid models.Grid, objects []models.Object) (int, int) {
	// naive approach, if more than 15 time & not found, stop.
	tryCount := 0
	var (
		x int
		y int
	)

	for {
		x, y = randomCoordinates(grid.Width, grid.Height)

		cell := grid.CellAtCoordinates(x, y)

		hasObjectAtCoordinate := false
		for _, object := range objects {
			if object.PositionX == x && object.PositionY == y {
				hasObjectAtCoordinate = true
			}
		}

		hasCharacterAtCoordinate := false
		for _, character := range characters {
			strategyHandler := NewNaiveRangeStrategyHandler(RangeCalculation{
				basePositionX:   x,
				basePositionY:   y,
				targetPositionX: character.PositionX,
				targetPositionY: character.PositionY,
				rangeLimit:      minimalRange,
			})

			if strategyHandler.PositionIsInRange() {
				hasCharacterAtCoordinate = true
				break
			}
		}

		fmt.Printf("generated : %v:%v - cell(%v),object(%v),character(%v)\n", x, y, cell.Type, hasObjectAtCoordinate, hasCharacterAtCoordinate)
		if cell.Type == models.CELL_TYPE_WALKABLE && !hasObjectAtCoordinate && !hasCharacterAtCoordinate {
			break
		}

		if tryCount == 15 {
			return 0, 0
		}
		tryCount++
	}

	return x, y
}

func getPlayerMinimalRange(playerCount int, gridHeight int, gridWidth int) int {
	return 1
}
