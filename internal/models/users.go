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

//func (u User) MarshalJSON() ([]byte, error) {
//	j, err := json.Marshal(struct {
//		Id string
//		Name string
//	}{
//		Id: fmt.,
//		Name: u.Name,
//	})
//
//	log.Printf("id : %s; name : %s\n", u._id, u.Name)
//
//	if err != nil {
//		return nil, err
//	}
//	return j, nil
//}
