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

func GenerateGrid(width int, height int) (*models.Grid, error) {
	return nil, nil
}

func randomCoordinates(width int, height int) (int, int) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(width), rand.Intn(height)
}
