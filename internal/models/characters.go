package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	User       primitive.ObjectID `json:"user,omitempty"`
	BloodSugar int                `json:"bloodSugar,omitempty"`
	Inventory  []Object           `json:"inventory,omitempty"`
	PositionX  int                `json:"positionX,omitempty"`
	PositionY  int                `json:"positionY,omitempty"`
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
