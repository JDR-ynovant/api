package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	BloodSugar int                `json:"bloodSugar"`
	Inventory  []Object           `json:"inventory"`
	PositionX  int                `json:"positionX"`
	PositionY  int                `json:"positionY"`
}

func (c Character) HasItem(id primitive.ObjectID) bool {
	hasItem := false
	for _, object := range c.Inventory {
		if object.Id == id {
			hasItem = true
		}
	}

	return hasItem
}
