package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"testing"
)

func TestHasObjectAtCoordinates(t *testing.T) {
	testObjects := [][]models.Object{
		{models.Object{PositionX: 10, PositionY: 5}, models.Object{PositionX: 7, PositionY: 6}},
		{models.Object{PositionX: 10, PositionY: 5}, models.Object{PositionX: 7, PositionY: 6}},
		{models.Object{PositionX: 10, PositionY: 5}, models.Object{PositionX: 7, PositionY: 6}},
	}
	testCoordinates := [][]int{{4, 5}, {8, 9}, {7, 6}}
	expected := []bool{false, false, true}

	for i := 0; i < 3; i++ {
		result := hasObjectAtCoordinates(testObjects[i], testCoordinates[i][0], testCoordinates[i][1])
		if result != expected[i] {
			t.Errorf("expected %v, got %v at iteration %v", expected[i], result, i)
		}
	}
}
