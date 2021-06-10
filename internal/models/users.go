package models

import (
	"github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	Subscription webpush.Subscription `json:"-"`
	Games        []primitive.ObjectID `json:"games,omitempty"`
}

func (u User) HasGameAttached(gameID string) bool {
	gameIDPrimitive, _ := primitive.ObjectIDFromHex(gameID)

	for _, objectID := range u.Games {
		if objectID == gameIDPrimitive {
			return true
		}
	}

	return false
}

func (u User) RemoveGame(game primitive.ObjectID) {
	var playerIndex int
	for i := 0; i < len(u.Games); i++ {
		if game == u.Games[i] {
			playerIndex = i
		}
	}

	u.Games[playerIndex] = u.Games[len(u.Games)-1]
	u.Games = u.Games[:len(u.Games)-1]
}
