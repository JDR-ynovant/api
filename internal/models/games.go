package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Game struct {
	Id             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name           string             `json:"name,omitempty"`
	Players        []Character        `json:"players,omitempty"`
	Playing        primitive.ObjectID `json:"playing,omitempty"`
	Owner          primitive.ObjectID `json:"owner,omitempty"`
	Status         GameStatus         `json:"status,omitempty"`
	Grid           Grid               `json:"grid,omitempty"`
	ExpiryDate     time.Time          `json:"expiryDate,omitempty"`
	Turns          []Turn             `json:"turns,omitempty"`
	TurnNumber     int                `json:"turnNumber,omitempty"`
	TurnExpiryDate time.Time          `json:"turnExpiryDate,omitempty"`
}

type GameStatus string

const (
	GAME_STATUS_CREATED  GameStatus = "created"
	GAME_STATUS_STARTED  GameStatus = "started"
	GAME_STATUS_FINISHED GameStatus = "finished"
)
