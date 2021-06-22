package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func getObjectIdFromString(s string) primitive.ObjectID {
	o, _ := primitive.ObjectIDFromHex(s)
	return o
}

func TestCalculateNewPlaying(t *testing.T) {
	testObjects := []models.Game{
		models.Game{
			// Player 1, 2 and 3 have played. 1 is dead, 3 just played so 2 must be selected
			Id: getObjectIdFromString("60d1c3f78638ce3c710d78e5"),
			Players: []models.Character{
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					BloodSugar: 10,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					BloodSugar: 2,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					BloodSugar: 4,
				},
			},
			PlayerCount: 3,
			Playing:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
			TurnNumber:  3,
			Turns: []models.Turn{
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca58"), // Turn 1
					Player:     getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					TurnNumber: 0,
				},
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca59"), // Turn 2
					Player:     getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					TurnNumber: 1,
				},
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca60"), // Turn 3
					Player:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					TurnNumber: 2,
				},
			},
		},
		models.Game{
			// Player 1, 2 and 3 have played. 1, 2 are dead, 3 just played but is selected
			Id: getObjectIdFromString("60d1c3f78638ce3c710d78e6"),
			Players: []models.Character{
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					BloodSugar: 10,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					BloodSugar: 10,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					BloodSugar: 4,
				},
			},
			PlayerCount: 3,
			Playing:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
			TurnNumber:  3,
			Turns: []models.Turn{
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca58"), // Turn 1
					Player:     getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					TurnNumber: 0,
				},
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca59"), // Turn 2
					Player:     getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					TurnNumber: 1,
				},
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca60"), // Turn 3
					Player:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					TurnNumber: 2,
				},
			},
		},
		models.Game{
			// Player 1, 2 and 3 have played. 4 has not played yet
			Id: getObjectIdFromString("60d1c3f78638ce3c710d78e6"),
			Players: []models.Character{
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					BloodSugar: 10,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					BloodSugar: 2,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					BloodSugar: 4,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb568"), // Player 4
					BloodSugar: 0,
				},
			},
			PlayerCount: 3,
			Playing:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
			TurnNumber:  3,
			Turns: []models.Turn{
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca58"), // Turn 1
					Player:     getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					TurnNumber: 0,
				},
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca59"), // Turn 2
					Player:     getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					TurnNumber: 1,
				},
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca60"), // Turn 3
					Player:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					TurnNumber: 2,
				},
			},
		},
		models.Game{
			// Player 1 has played. 2 is the next
			Id: getObjectIdFromString("60d1c3f78638ce3c710d78e6"),
			Players: []models.Character{
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					BloodSugar: 0,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
					BloodSugar: 0,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
					BloodSugar: 0,
				},
				models.Character{
					Id:         getObjectIdFromString("60d1b604117863009d4eb568"), // Player 4
					BloodSugar: 0,
				},
			},
			PlayerCount: 3,
			Playing:     getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
			TurnNumber:  3,
			Turns: []models.Turn{
				models.Turn{
					Id:         getObjectIdFromString("60d1b772a4d26546f082ca58"), // Turn 1
					Player:     getObjectIdFromString("60d1b604117863009d4eb565"), // Player 1
					TurnNumber: 0,
				},
			},
		},
	}

	expected := []primitive.ObjectID{
		getObjectIdFromString("60d1b604117863009d4eb566"), // Player 1
		getObjectIdFromString("60d1b604117863009d4eb567"), // Player 3
		getObjectIdFromString("60d1b604117863009d4eb568"), // Player 4
		getObjectIdFromString("60d1b604117863009d4eb566"), // Player 2
	}

	for i := 0; i < len(testObjects); i++ {
		result := calculateNewPlaying(&testObjects[i])
		if result.Hex() != expected[i].Hex() {
			t.Errorf("expected %v, got %v at iteration %v", expected[i], result, i)
		}
	}
}
