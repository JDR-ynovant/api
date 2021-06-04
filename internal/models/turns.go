package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Turn struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Actions    []Action           `json:"actions,omitempty"`
	X          int                `json:"x,omitempty"`
	Y          int                `json:"y,omitempty"`
	Player     primitive.ObjectID `json:"player,omitempty"`
	TurnNumber int                `json:"turnNumber,omitempty"`
}

type Action struct {
	Type string `json:"type,omitempty"`
}

type MoveAction struct {
	Action
	TargetX int `json:"targetX,omitempty"`
	TargetY int `json:"targetY,omitempty"`
}

type AttackAction struct {
	MoveAction
	Character primitive.ObjectID `json:"character,omitempty"`
	Object    primitive.ObjectID `json:"weapon,omitempty"`
}

type UsageAction struct {
	Action
	Object primitive.ObjectID `json:"object,omitempty"`
}

type NullAction struct {
	Action
	Reason string `json:"reason,omitempty"`
}
