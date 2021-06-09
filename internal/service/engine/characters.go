package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CHARACTER_MAX_BLOOD_SUGAR = 15

func CreateCharacter(user primitive.ObjectID) *models.Character {
	character := models.Character{
		Id:         primitive.NewObjectID(),
		User:       user,
		BloodSugar: CHARACTER_MAX_BLOOD_SUGAR,
		Inventory:  make([]models.Object, 0),
		PositionX:  0,
		PositionY:  0,
	}

	return &character
}
