package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"math/rand"
	"time"
)

func GenerateObjects(grid *models.Grid, playerCount int) []models.Object {
	objects := make([]models.Object, 0)
	count := int(math.Round(BASE_OBJECT_COUNT * float64(playerCount)))

	for i := 0; i < count; i++ {
		var (
			x int
			y int
		)

		for {
			x, y = randomCoordinates(grid.Width, grid.Height)
			cell := grid.CellAtCoordinates(x, y)
			cellIsNotOnLimitBound := x != 0 && y != 0 && x != grid.Width && y != grid.Height

			if cellIsNotOnLimitBound && cell.Type != models.CELL_TYPE_OBSTACLE && !hasObjectAtCoordinates(objects, x, y) {
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
		Id:        primitive.NewObjectID(),
		Name:      randomFromArray(getAvailableObjectNames()),
		Type:      randomFromArrayObjectType(getAvailableObjectTypes()),
		Value:     randomObjectValue(),
		PositionX: x,
		PositionY: y,
		Picked:    false,
	}
}

func getAvailableObjectNames() []string {
	return []string{"Sugar Cane", "Candy", "Spear of sugar", "Liquid Sugar"}
}

func getAvailableObjectTypes() []models.ObjectType {
	return []models.ObjectType{models.OBJECT_TYPE_WEAPON, models.OBJECT_TYPE_POTION}
}

func randomFromArray(strings []string) string {
	return strings[rand.Intn(len(strings))]
}

func randomFromArrayObjectType(objectTypes []models.ObjectType) models.ObjectType {
	return objectTypes[rand.Intn(len(objectTypes))]
}

func randomObjectValue() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(models.OBJECT_MAXIMAL_VALUE)
}
