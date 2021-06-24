package models

import (
	"github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	Subscription webpush.Subscription `json:"-"`
	Games        []MetaGame           `json:"games"`
}

type MetaGame struct {
	Id     primitive.ObjectID
	Name   string
	Status GameStatus
}

func MetaFromGame(game Game) MetaGame {
	return MetaGame{
		Id:     game.Id,
		Name:   game.Name,
		Status: game.Status,
	}
}

func (u User) GetGame(gameId primitive.ObjectID) *MetaGame {
	for i, game := range u.Games {
		if game.Id == gameId {
			return &u.Games[i]
		}
	}
	return nil
}

func (u User) HasGameAttached(gameID string) bool {
	gameIDPrimitive, _ := primitive.ObjectIDFromHex(gameID)

	for _, game := range u.Games {
		if game.Id == gameIDPrimitive {
			return true
		}
	}

	return false
}

func (u *User) RemoveGame(game primitive.ObjectID) {
	var playerIndex int
	for i := 0; i < len(u.Games); i++ {
		if game == u.Games[i].Id {
			playerIndex = i
		}
	}

	u.Games[playerIndex] = u.Games[len(u.Games)-1]
	u.Games = u.Games[:len(u.Games)-1]
}
