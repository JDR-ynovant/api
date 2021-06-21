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

type ActionType string

const (
	ACTION_TYPE_MOVE   ActionType = "move"
	ACTION_TYPE_ATTACK ActionType = "attack"
	ACTION_TYPE_USAGE  ActionType = "usage"
	ACTION_TYPE_NULL   ActionType = "null"
)

type Action struct {
	Type ActionType `json:"type,omitempty"`

	// MoveAction
	TargetX int `json:"targetX,omitempty"`
	TargetY int `json:"targetY,omitempty"`
	// AttackAction & UsageAction
	Character primitive.ObjectID `json:"character,omitempty"`
	Object    primitive.ObjectID `json:"weapon,omitempty"`
	// NullAction
	Reason string `json:"reason,omitempty"`
}

func (a Action) Validate() bool {
	if !isValidActionType(a.Type) {
		return false
	}

	switch a.Type {
	case ACTION_TYPE_USAGE:
	case ACTION_TYPE_ATTACK:
		return a.Character != primitive.ObjectID{} && a.Object != primitive.ObjectID{}
	case ACTION_TYPE_MOVE:
		return true
	case ACTION_TYPE_NULL:
		return a.Reason != ""
	}

	return false
}

func isValidActionType(actionType ActionType) bool {
	switch actionType {
	case ACTION_TYPE_MOVE:
	case ACTION_TYPE_ATTACK:
	case ACTION_TYPE_USAGE:
	case ACTION_TYPE_NULL:
		return true
	}
	return false
}
