package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Grid struct {
	Id     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Width  int                `json:"width,omitempty"`
	Height int                `json:"height,omitempty"`
	Cells  []Cell             `json:"cells,omitempty"`
}

func (g Grid) CellAtCoordinates(x int, y int) *Cell {
	return nil
}
