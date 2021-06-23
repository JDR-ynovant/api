package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestMoveActionHandler_IsLegit(t *testing.T) {
	gridPattern := [][][]int{
		{
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{1, 0, 0, 1, 0},
		},
	}

	testObjects := []testObjectsStruct{
		{
			r: RangeCalculation{
				grid: &models.Grid{
					Id:     primitive.ObjectID{},
					Width:  5,
					Height: 5,
					Cells:  generateFromGridPattern(gridPattern[0]),
				},
				basePositionX:   2,
				basePositionY:   2,
				targetPositionX: 1,
				targetPositionY: 1,
				rangeLimit:      1,
			},
			s: STRATEGY_RANGE,
		},
		{
			r: RangeCalculation{
				grid: &models.Grid{
					Id:     primitive.ObjectID{},
					Width:  5,
					Height: 5,
					Cells:  generateFromGridPattern(gridPattern[0]),
				},
				basePositionX:   2,
				basePositionY:   2,
				targetPositionX: 2,
				targetPositionY: 3,
				rangeLimit:      1,
			},
			s: STRATEGY_RANGE,
		},
		{
			r: RangeCalculation{
				grid: &models.Grid{
					Id:     primitive.ObjectID{},
					Width:  5,
					Height: 5,
					Cells:  generateFromGridPattern(gridPattern[0]),
				},
				basePositionX:   2,
				basePositionY:   2,
				targetPositionX: 2,
				targetPositionY: 4,
				rangeLimit:      1,
			},
			s: STRATEGY_RANGE,
		},
		{
			r: RangeCalculation{
				grid: &models.Grid{
					Id:     primitive.ObjectID{},
					Width:  5,
					Height: 5,
					Cells:  generateFromGridPattern(gridPattern[0]),
				},
				basePositionX:   2,
				basePositionY:   4,
				targetPositionX: 2,
				targetPositionY: 5,
				rangeLimit:      1,
			},
			s: STRATEGY_RANGE,
		},
	}

	expectedOutput := []bool{
		true,
		true,
		false,
		true,
	}

	for i, testObject := range testObjects {
		handler := NewNaiveRangeStrategyHandler(testObject.r)
		if result := handler.PositionIsInRange(); result != expectedOutput[i] {
			t.Errorf("expected %v, got %v at iteration %v", expectedOutput[i], result, i)
		}
	}
}
