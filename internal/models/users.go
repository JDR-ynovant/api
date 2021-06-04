package models

import (
	"github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	Subscription webpush.Subscription `json:"-"`
}
