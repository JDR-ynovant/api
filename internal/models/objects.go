package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Object struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Type      string             `json:"type,omitempty"`
	Value     int                `json:"value,omitempty"`
	PositionX int                `json:"positionX,omitempty"`
	PositionY int                `json:"positionY,omitempty"`
	Picked    bool               `json:"picked,omitempty"`
}

//type Weapon struct {
//	Object
//	Damage int `json:"damage,omitempty"`
//}
//
//type Potion struct {
//	Object
//	Value int        `json:"value,omitempty"`
//	Kind  PotionType `json:"kind,omitempty"`
//}
//
//type PotionType string
//
//const (
//	POTION_TYPE_DAMAGE PotionType = "damage"
//	POTION_TYPE_HEAL   PotionType = "HEAL"
//)
