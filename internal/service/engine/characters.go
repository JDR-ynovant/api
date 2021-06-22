package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCharacter(user primitive.ObjectID) *models.Character {
	character := models.Character{
		Id:         primitive.NewObjectID(),
		User:       user,
		BloodSugar: 0,
		Inventory:  make([]models.Object, 0),
		PositionX:  0,
		PositionY:  0,
	}

	return &character
}
