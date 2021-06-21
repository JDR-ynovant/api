package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Object struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Type      ObjectType         `json:"type,omitempty"`
	Value     int                `json:"value,omitempty"`
	PositionX int                `json:"positionX,omitempty"`
	PositionY int                `json:"positionY,omitempty"`
	Picked    bool               `json:"picked,omitempty"`
}

const (
	OBJECT_MINIMAL_VALUE = 1
	OBJECT_MAXIMAL_VALUE = 4
)

type ObjectType string

const (
	OBJECT_TYPE_WEAPON ObjectType = "weapon"
	OBJECT_TYPE_POTION ObjectType = "potion"
)
