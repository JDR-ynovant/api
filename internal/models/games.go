package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Game struct {
	Id             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name           string             `json:"name,omitempty"`
	Players        []Character        `json:"players"`
	PlayerCount    int                `json:"playerCount"`
	Playing        primitive.ObjectID `json:"playing,omitempty"`
	Owner          primitive.ObjectID `json:"owner,omitempty"`
	Status         GameStatus         `json:"status,omitempty"`
	Grid           primitive.ObjectID `json:"grid,omitempty"`
	Objects        []Object           `json:"objects"`
	ExpiryDate     time.Time          `json:"expiryDate,omitempty"`
	Turns          []Turn             `json:"turns"`
	TurnNumber     int                `json:"turnNumber"`
	TurnExpiryDate time.Time          `json:"turnExpiryDate,omitempty"`
}

type GameStatus string

const (
	GAME_STATUS_CREATED  GameStatus = "created"
	GAME_STATUS_STARTED  GameStatus = "started"
	GAME_STATUS_FINISHED GameStatus = "finished"
)

func (g Game) GetPlayer(userId primitive.ObjectID) *Character {
	for _, player := range g.Players {
		if player.Id == userId {
			return &player
		}
	}
	return nil
}

func (g Game) GetObject(objectID primitive.ObjectID) *Object {
	for _, object := range g.Objects {
		if object.Id == objectID {
			return &object
		}
	}
	return nil
}

func (g Game) HasPlayer(userId primitive.ObjectID) bool {
	for _, player := range g.Players {
		if player.Id == userId {
			return true
		}
	}
	return false
}

func (g *Game) RemovePlayer(userId primitive.ObjectID) {
	var playerIndex int
	for i := 0; i < len(g.Players); i++ {
		if userId == g.Players[i].Id {
			playerIndex = i
		}
	}

	g.Players[playerIndex] = g.Players[len(g.Players)-1]
	g.Players = g.Players[:len(g.Players)-1]
}
