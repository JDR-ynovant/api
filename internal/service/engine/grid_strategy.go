package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
)

type Strategy string

const (
	STRATEGY_RANGE      Strategy = "range"
	STRATEGY_GRID_RANGE Strategy = "grid_range"
	STRATEGY_PATH       Strategy = "path"
)

type GridStrategyHandler interface {
	PositionIsInRange() bool
}

func GetGridStrategyHandler(strategy Strategy, r RangeCalculation) GridStrategyHandler {
	switch strategy {
	case STRATEGY_PATH:
		return NewPathStrategyHandler(r)
	case STRATEGY_RANGE:
		return NewNaiveRangeStrategyHandler(r)
	case STRATEGY_GRID_RANGE:
		return NewNaiveRangeStrategyHandler(r)
	}
	return nil
}

// =================== NaiveRangeStrategyHandler

type NaiveRangeStrategyHandler struct {
	basePositionX   int
	basePositionY   int
	targetPositionX int
	targetPositionY int
	rangeLimit      int
}

func NewNaiveRangeStrategyHandler(r RangeCalculation) *NaiveRangeStrategyHandler {
	return &NaiveRangeStrategyHandler{
		basePositionX:   r.basePositionX,
		basePositionY:   r.basePositionY,
		targetPositionX: r.targetPositionX,
		targetPositionY: r.targetPositionY,
		rangeLimit:      r.rangeLimit}
}

func (r NaiveRangeStrategyHandler) PositionIsInRange() bool {
	// basic circle around position check
	return (r.basePositionX-r.rangeLimit <= r.targetPositionX && r.targetPositionX <= r.basePositionX+r.rangeLimit) &&
		(r.basePositionY-r.rangeLimit <= r.targetPositionY && r.targetPositionY <= r.basePositionY+r.rangeLimit)
}

// =================== PathStrategyHandler

type GridRangeStrategyHandler struct {
	grid            *models.Grid
	basePositionX   int
	basePositionY   int
	targetPositionX int
	targetPositionY int
	rangeLimit      int
}

func NewGridRangeStrategyHandler(r RangeCalculation) *GridRangeStrategyHandler {
	return &GridRangeStrategyHandler{
		grid:            r.grid,
		basePositionX:   r.basePositionX,
		basePositionY:   r.basePositionY,
		targetPositionX: r.targetPositionX,
		targetPositionY: r.targetPositionY,
		rangeLimit:      r.rangeLimit}
}

func (g GridRangeStrategyHandler) PositionIsInRange() bool {
	// basic circle around position check with grid range check
	return (g.basePositionX-g.rangeLimit <= g.targetPositionX && g.targetPositionX <= g.basePositionX+g.rangeLimit) &&
		(g.basePositionY-g.rangeLimit <= g.targetPositionY && g.targetPositionY <= g.basePositionY+g.rangeLimit) &&
		(g.targetPositionX > 0 && g.targetPositionX < g.grid.Height) &&
		(g.targetPositionY > 0 && g.targetPositionY < g.grid.Width)
}

// =================== PathStrategyHandler

func NewPathStrategyHandler(r RangeCalculation) *PathStrategyHandler {
	return &PathStrategyHandler{
		grid:            r.grid,
		basePositionX:   r.basePositionX,
		basePositionY:   r.basePositionY,
		targetPositionX: r.targetPositionX,
		targetPositionY: r.targetPositionY,
		rangeLimit:      r.rangeLimit}
}

type PathStrategyHandler struct {
	grid            *models.Grid
	basePositionX   int
	basePositionY   int
	targetPositionX int
	targetPositionY int
	rangeLimit      int
}

func (p PathStrategyHandler) PositionIsInRange() bool {
	panic("implement me")
	return true
}
