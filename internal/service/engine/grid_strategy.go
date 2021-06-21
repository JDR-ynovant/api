package engine

type Strategy string

const (
	STRATEGY_RANGE Strategy = "range"
	STRATEGY_PATH  Strategy = "path"
)

type GridStrategyHandler interface {
	PositionIsInRange() bool
}

func GetGridStrategyHandler(strategy Strategy, r RangeCalculation) GridStrategyHandler {
	switch strategy {
	case STRATEGY_PATH:
		return NewPathStrategyHandler(r.basePositionX, r.basePositionY, r.targetPositionX, r.basePositionY, r.rangeLimit)
	case STRATEGY_RANGE:
		return NewNaiveRangeStrategyHandler(r.basePositionX, r.basePositionY, r.targetPositionX, r.basePositionY, r.rangeLimit)
	}
	return nil
}

type NaiveRangeStrategyHandler struct {
	basePositionX   int
	basePositionY   int
	targetPositionX int
	targetPositionY int
	rangeLimit      int
}

func NewNaiveRangeStrategyHandler(basePositionX int, basePositionY int, targetPositionX int, targetPositionY int, rangeLimit int) *NaiveRangeStrategyHandler {
	return &NaiveRangeStrategyHandler{
		basePositionX:   basePositionX,
		basePositionY:   basePositionY,
		targetPositionX: targetPositionX,
		targetPositionY: targetPositionY,
		rangeLimit:      rangeLimit}
}

func (r NaiveRangeStrategyHandler) PositionIsInRange() bool {
	// basic circle around position check
	return (r.basePositionX-r.rangeLimit <= r.targetPositionX || r.targetPositionX <= r.basePositionX+r.rangeLimit) &&
		(r.basePositionY-r.rangeLimit <= r.targetPositionY || r.targetPositionY <= r.basePositionY+r.rangeLimit)
}

func NewPathStrategyHandler(basePositionX int, basePositionY int, targetPositionX int, targetPositionY int, rangeLimit int) *PathStrategyHandler {
	return &PathStrategyHandler{
		basePositionX:   basePositionX,
		basePositionY:   basePositionY,
		targetPositionX: targetPositionX,
		targetPositionY: targetPositionY,
		rangeLimit:      rangeLimit}
}

type PathStrategyHandler struct {
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
