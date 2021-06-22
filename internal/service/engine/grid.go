package engine

import (
	"github.com/JDR-ynovant/api/internal"
	"github.com/JDR-ynovant/api/internal/models"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"time"
)

const (
	DEFAULT_GRID_WIDTH  = 15
	DEFAULT_GRID_HEIGHT = 15
)

func GenerateGrid(width int, height int) *models.Grid {
	cellMapping := map[models.CellType]string{
		models.CELL_TYPE_OBSTACLE: "/assets/img/obstacle.png",
		models.CELL_TYPE_WALKABLE: "/assets/img/grass.png",
	}

	grid := models.Grid{
		Width:  width,
		Height: height,
		Cells:  make([]models.Cell, 0),
	}

	for currentX := 0; currentX < width; currentX++ {
		for currentY := 0; currentY < height; currentY++ {
			cellType := randomCellType()
			cell := models.Cell{
				X: currentX,
				Y: currentY,
				// @todo randomize
				Type:   cellType,
				Sprite: cellMapping[cellType],
			}

			grid.Cells = append(grid.Cells, cell)
		}
	}

	return &grid
}

func randomCellType() models.CellType {
	config := internal.GetConfig()

	rand.Seed(time.Now().UnixNano())
	chooser, _ := wr.NewChooser(
		wr.Choice{Item: models.CELL_TYPE_OBSTACLE, Weight: config.RuleObstacleQuota},
		wr.Choice{Item: models.CELL_TYPE_WALKABLE, Weight: config.RuleWalkableQuota},
	)

	return chooser.Pick().(models.CellType)
}

func randomCoordinates(width int, height int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(width), rand.Intn(height)
}

type RangeCalculation struct {
	grid            *models.Grid
	basePositionX   int
	basePositionY   int
	targetPositionX int
	targetPositionY int
	rangeLimit      int
}

func IsInRange(r RangeCalculation, strategy Strategy) bool {
	handler := GetGridStrategyHandler(strategy, r)
	return handler.PositionIsInRange()
}
