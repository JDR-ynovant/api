package models

type Cell struct {
	X      int      `json:"x"`
	Y      int      `json:"y"`
	Type   CellType `json:"type,omitempty"`
	Sprite string   `json:"sprite,omitempty"`
}

type CellType string

const (
	CELL_TYPE_WALKABLE CellType = "walkable"
	CELL_TYPE_OBSTACLE CellType = "obstacle"
)
