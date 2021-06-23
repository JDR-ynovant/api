package engine

import (
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type actionsTestObjectStruct struct {
	game   models.Game
	grid   models.Grid
	action models.Action
	turn   models.Turn
}

type actionsTestExpected struct {
	result bool
	error  string
}

func TestMoveActionHandler_IsLegit(t *testing.T) {
	gridPattern := [][][]int{
		{
			{1, 0, 1, 0, 0, 0, 1},
			{0, 0, 1, 0, 0, 0, 1},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1, 0, 1},
			{0, 1, 0, 0, 1, 0, 0},
		},
	}

	testObjects := []actionsTestObjectStruct{
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						PositionX: 3,
						PositionY: 3,
					},
				},
				Objects:    nil,
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:    models.ACTION_TYPE_MOVE,
				TargetX: 2,
				TargetY: 3,
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						PositionX: 2,
						PositionY: 2,
					},
				},
				Objects:    nil,
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:    models.ACTION_TYPE_MOVE,
				TargetX: 6,
				TargetY: 6,
			},
			turn: models.Turn{
				X:          2,
				Y:          2,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						PositionX: 3,
						PositionY: 3,
					},
				},
				Objects:    nil,
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:    models.ACTION_TYPE_MOVE,
				TargetX: 4,
				TargetY: 2,
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						PositionX: 3,
						PositionY: 3,
					},
				},
				Objects:    nil,
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:    models.ACTION_TYPE_MOVE,
				TargetX: 5,
				TargetY: 2,
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						PositionX: 6,
						PositionY: 6,
					},
				},
				Objects:    nil,
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:    models.ACTION_TYPE_MOVE,
				TargetX: 7,
				TargetY: 6,
			},
			turn: models.Turn{
				X:          6,
				Y:          6,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
	}

	expectedOutput := []actionsTestExpected{
		{
			result: false, // non legit obstacle
			error:  "move - out of range or not walkable",
		},
		{
			result: false, // non legit out of range
			error:  "move - out of range or not walkable",
		},
		{
			result: true, // legit 1 range
			error:  "",
		},
		{
			result: true, // legit 2 range
			error:  "",
		},
		{
			result: false, // non legit out of grid
			error:  "move - out of range or not walkable",
		},
	}

	for i, testObject := range testObjects {
		moveActionHandler := GetActionHandler(&testObject.game, &testObject.grid, &testObject.action, &testObject.turn)

		if result, err := moveActionHandler.IsLegit(); result != expectedOutput[i].result || err != expectedOutput[i].error {
			t.Errorf("expected %v, got %v at iteration %v", expectedOutput[i], actionsTestExpected{result, err}, i)
		}
	}
}

func TestAttackActionHandler_IsLegit(t *testing.T) {
	gridPattern := [][][]int{
		{
			{1, 0, 1, 0, 0, 0, 1},
			{0, 0, 1, 0, 0, 0, 1},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1, 0, 1},
			{0, 1, 0, 0, 1, 0, 0},
		},
	}

	testObjects := []actionsTestObjectStruct{
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						PositionX: 3,
						PositionY: 3,
					},
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb566"),
						PositionX: 2,
						PositionY: 3,
					},
				},
				Objects:    nil,
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:      models.ACTION_TYPE_ATTACK,
				Character: getObjectIdFromString("60d1b604117863009d4eb566"),
				Object:    primitive.ObjectID{},
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id: getObjectIdFromString("60d1b604117863009d4eb565"),
						Inventory: []models.Object{
							{Id: getObjectIdFromString("60d1c3f78638ce3c710d78e5"), Picked: true},
						},
						PositionX: 3,
						PositionY: 3,
					},
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb566"),
						PositionX: 2,
						PositionY: 3,
					},
				},
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:      models.ACTION_TYPE_ATTACK,
				Character: getObjectIdFromString("60d1b604117863009d4eb566"),
				Object:    getObjectIdFromString("60d1c3f78638ce3c710d78e5"),
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id: getObjectIdFromString("60d1b604117863009d4eb565"),
						Inventory: []models.Object{
							{Id: getObjectIdFromString("60d1c3f78638ce3c710d78e5"), Picked: true},
						},
						PositionX: 3,
						PositionY: 3,
					},
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb566"),
						PositionX: 2,
						PositionY: 1,
					},
				},
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:      models.ACTION_TYPE_ATTACK,
				Character: getObjectIdFromString("60d1b604117863009d4eb566"),
				Object:    getObjectIdFromString("60d1c3f78638ce3c710d78e5"),
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
		{
			game: models.Game{
				Players: []models.Character{
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb565"),
						Inventory: []models.Object{},
						PositionX: 3,
						PositionY: 3,
					},
					models.Character{
						Id:        getObjectIdFromString("60d1b604117863009d4eb566"),
						PositionX: 2,
						PositionY: 3,
					},
				},
				Turns:      make([]models.Turn, 0),
				TurnNumber: 0,
			},
			grid: models.Grid{
				Width:  len(gridPattern[0][0]),
				Height: len(gridPattern[0]),
				Cells:  generateFromGridPattern(gridPattern[0]),
			},
			action: models.Action{
				Type:      models.ACTION_TYPE_ATTACK,
				Character: getObjectIdFromString("60d1b604117863009d4eb566"),
				Object:    getObjectIdFromString("60d1c3f78638ce3c710d78e5"),
			},
			turn: models.Turn{
				X:          3,
				Y:          3,
				Player:     getObjectIdFromString("60d1b604117863009d4eb565"),
				TurnNumber: 0,
			},
		},
	}

	expectedOutput := []actionsTestExpected{
		{
			result: true, // legit no object
			error:  "",
		},
		{
			result: true, // legit with object
			error:  "",
		},
		{
			result: false, // non legit out of range
			error:  "attack - out of range or do not possess object",
		},
		{
			result: false, // non legit dont possess object
			error:  "attack - out of range or do not possess object",
		},
	}

	for i, testObject := range testObjects {
		moveActionHandler := GetActionHandler(&testObject.game, &testObject.grid, &testObject.action, &testObject.turn)

		if result, err := moveActionHandler.IsLegit(); result != expectedOutput[i].result || err != expectedOutput[i].error {
			t.Errorf("expected %v, got %v at iteration %v", expectedOutput[i], actionsTestExpected{result, err}, i)
		}
	}
}
