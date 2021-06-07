package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"math/rand"
	"time"
)

func GenerateObjects(grid *models.Grid, count int) []models.Object {
	objects := make([]models.Object, 0)

	for i := 0; i < count; i++ {
		var (
			x int
			y int
		)

		for {
			x, y = randomCoordinates(grid.Width, grid.Height)
			cell := grid.CellAtCoordinates(x, y)

			if cell.Type != models.CELL_TYPE_OBSTACLE && !hasObjectAtCoordinates(objects, x, y) {
				break
			}
		}

		objects = append(objects, randomObject(x, y))
	}

	return objects
}

func hasObjectAtCoordinates(objects []models.Object, x int, y int) bool {
	for _, object := range objects {
		if object.PositionX == x && object.PositionY == y {
			return true
		}
	}
	return false
}

func randomObject(x int, y int) models.Object {
	return models.Object{
		Name:      randomFromArray(getAvailableObjectNames()),
		Type:      randomFromArray(getAvailableObjectTypes()),
		Value:     randomObjectValue(),
		PositionX: x,
		PositionY: y,
		Picked:    false,
	}
}

func getAvailableObjectNames() []string {
	return []string{"Sugar Cane", "Candy", "Spear of sugar", "Liquid Sugar"}
}

func getAvailableObjectTypes() []string {
	return []string{"weapon", "potion"}
}

func randomFromArray(strings []string) string {
	return strings[rand.Intn(len(strings))]
}

func randomObjectValue() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(models.OBJECT_MAXIMAL_VALUE)
}
