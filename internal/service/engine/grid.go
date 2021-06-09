package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"math/rand"
	"time"
)

const (
	DEFAULT_GRID_WIDTH  = 15
	DEFAULT_GRID_HEIGHT = 15
)

func GenerateGrid(width int, height int) *models.Grid {
	grid := models.Grid{
		Width:  width,
		Height: height,
		Cells:  make([]models.Cell, 0),
	}

	for currentX := 0; currentX < width; currentX++ {
		for currentY := 0; currentY < height; currentY++ {
			cell := models.Cell{
				X:      currentX,
				Y:      currentY,
				Type:   models.CELL_TYPE_WALKABLE,
				Sprite: "/assets/img/grass.png",
			}

			grid.Cells = append(grid.Cells, cell)
		}
	}

	return &grid
}

func randomCoordinates(width int, height int) (int, int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(width), rand.Intn(height)
}
