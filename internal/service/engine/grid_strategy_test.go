package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type gridStrategyTestObjectsStruct struct {
	r RangeCalculation
	s Strategy
}

const (
	Walkable = 0
	Obstacle = 1
)

func TestNaiveRangeStrategyHandler_PositionIsInRange(t *testing.T) {
	gridPattern := [][][]int{
		{
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{1, 0, 0, 1, 0},
		},
	}

	testObjects := []gridStrategyTestObjectsStruct{
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

func TestGridRangeStrategyHandler_PositionIsInRange(t *testing.T) {
	gridPattern := [][][]int{
		{
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{1, 0, 0, 1, 0},
		},
	}

	testObjects := []gridStrategyTestObjectsStruct{
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
				targetPositionY: 4,
				rangeLimit:      2,
			},
			s: STRATEGY_RANGE,
		},
	}

	expectedOutput := []bool{
		true,
		true,
		false,
		false,
		true,
	}

	for i, testObject := range testObjects {
		handler := NewGridRangeStrategyHandler(testObject.r)
		if result := handler.PositionIsInRange(); result != expectedOutput[i] {
			t.Errorf("expected %v, got %v at iteration %v", expectedOutput[i], result, i)
		}
	}
}

func generateFromGridPattern(pattern [][]int) []models.Cell {
	cells := make([]models.Cell, 0)
	for xIndex, xLine := range pattern {
		for yIndex, yValue := range xLine {
			var cellType models.CellType
			if yValue == Walkable {
				cellType = models.CELL_TYPE_WALKABLE
			} else if yValue == Obstacle {
				cellType = models.CELL_TYPE_OBSTACLE
			}

			cells = append(cells, models.Cell{
				X:    xIndex,
				Y:    yIndex,
				Type: cellType,
			})
		}
	}

	return cells
}
