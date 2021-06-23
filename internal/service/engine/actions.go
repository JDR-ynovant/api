package engine

import (
	"errors"
	"github.com/JDR-ynovant/api/internal"
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActionHandler interface {
	Handle()
	IsLegit() (bool, error)
}

func GetActionHandler(game *models.Game, action *models.Action, turn *models.Turn) ActionHandler {
	switch action.Type {
	case models.ACTION_TYPE_MOVE:
		return NewMoveActionHandler(game, turn, action)
	case models.ACTION_TYPE_ATTACK:
		return NewAttackActionHandler(game, turn, action)
	case models.ACTION_TYPE_USAGE:
		return NewUsageActionHandler(game, turn, action)
	case models.ACTION_TYPE_NULL:
		return NewNullActionHandler(game, turn, action)
	}
	return nil
}

// ================== MoveActionHandler

// MoveActionHandler godoc
// A MoveAction action is legit if :
// - move is in range
// - move do not overlap with non walkable tile (obstacle, or player)
// An item is picked up if :
// - end coordinate overlap with object
type MoveActionHandler struct {
	game   *models.Game
	turn   *models.Turn
	action *models.Action
}

func NewMoveActionHandler(game *models.Game, turn *models.Turn, action *models.Action) *MoveActionHandler {
	return &MoveActionHandler{game: game, turn: turn, action: action}
}

func (m MoveActionHandler) Handle() {
	player := m.game.GetPlayer(m.turn.Player)
	player.PositionX = m.action.TargetX
	player.PositionY = m.action.TargetY

	m.game.Turns[m.game.TurnNumber].Actions = append(m.game.Turns[m.game.TurnNumber].Actions, *m.action)
}

func (m MoveActionHandler) IsLegit() (bool, error) {
	config := internal.GetConfig()
	player := m.game.GetPlayer(m.turn.Player)

	gr := repository.NewGridRepository()
	grid, _ := gr.FindOneById(m.game.Grid.Hex())

	r := RangeCalculation{
		basePositionX:   player.PositionX,
		basePositionY:   player.PositionY,
		targetPositionX: m.action.TargetX,
		targetPositionY: m.action.TargetY,
		rangeLimit:      config.RuleAttackRange,
	}

	targetCell := grid.CellAtCoordinates(m.action.TargetX, m.action.TargetY)

	if IsInRange(r, STRATEGY_GRID_RANGE) && targetCell.Type == models.CELL_TYPE_WALKABLE {
		return true, nil
	}

	return false, errors.New("move - out of range or not walkable")
}

// ================== AttackActionHandler

// AttackActionHandler godoc
// A AttackAction action is legit if :
// - player and target are in range
// - player has this weapon in inventory (if weapon is null, base damage is used)
type AttackActionHandler struct {
	game   *models.Game
	turn   *models.Turn
	action *models.Action
}

func NewAttackActionHandler(game *models.Game, turn *models.Turn, action *models.Action) *AttackActionHandler {
	return &AttackActionHandler{game: game, turn: turn, action: action}
}

func (a AttackActionHandler) Handle() {
	config := internal.GetConfig()
	target := a.game.GetPlayer(a.action.Character)

	if target.BloodSugar < config.RuleBloodSugarCap {
		// if object is empty, apply base damage
		emptyObj := primitive.ObjectID{}
		if a.action.Object == emptyObj {
			target.BloodSugar += config.RuleBaseDamage
		} else {
			// else apply weapon damage
			object := a.game.GetObject(a.action.Object)

			if object.Type == models.OBJECT_TYPE_WEAPON {
				target.BloodSugar += object.Value
			}
		}

		if target.BloodSugar >= config.RuleBloodSugarCap {
			target.BloodSugar = -1
		}
	}

	a.game.Turns[a.game.TurnNumber].Actions = append(a.game.Turns[a.game.TurnNumber].Actions, *a.action)
}

func (a AttackActionHandler) IsLegit() (bool, error) {
	config := internal.GetConfig()
	target := a.game.GetPlayer(a.action.Character)
	player := a.game.GetPlayer(a.turn.Player)

	r := RangeCalculation{
		basePositionX:   player.PositionX,
		basePositionY:   player.PositionY,
		targetPositionX: target.PositionX,
		targetPositionY: target.PositionY,
		rangeLimit:      config.RuleAttackRange,
	}

	if IsInRange(r, STRATEGY_RANGE) && player.HasItem(a.action.Object) {
		return true, nil
	}
	return false, errors.New("attack - out of range or do not possess object")
}

// ================== UsageActionHandler

// UsageActionHandler godoc
// A UsageAction action is legit if :
// - player have this item in his inventory
// Player blood sugar killed this turn is set to -1, an reset later (for detection)
type UsageActionHandler struct {
	game   *models.Game
	turn   *models.Turn
	action *models.Action
}

func NewUsageActionHandler(game *models.Game, turn *models.Turn, action *models.Action) *UsageActionHandler {
	return &UsageActionHandler{game: game, turn: turn, action: action}
}

func (u UsageActionHandler) Handle() {
	object := u.game.GetObject(u.action.Object)
	target := u.game.GetPlayer(u.action.Character)

	if object.Type == models.OBJECT_TYPE_POTION && target.BloodSugar > 0 {
		target.BloodSugar -= object.Value
	}

	u.game.Turns[u.game.TurnNumber].Actions = append(u.game.Turns[u.game.TurnNumber].Actions, *u.action)
}

func (u UsageActionHandler) IsLegit() (bool, error) {
	hasItem := false
	for _, object := range u.game.GetPlayer(u.turn.Player).Inventory {
		if object.Id == u.action.Object {
			hasItem = true
		}
	}

	if hasItem {
		return true, nil
	}
	return false, errors.New("use - object not in possession")
}

// ================== NullActionHandler

// NullActionHandler godoc
// A NullAction action is legit if :
// - always, a turn can always be skip. but this is the last action of the turn
type NullActionHandler struct {
	game   *models.Game
	turn   *models.Turn
	action *models.Action
}

func NewNullActionHandler(game *models.Game, turn *models.Turn, action *models.Action) *NullActionHandler {
	return &NullActionHandler{game: game, turn: turn, action: action}
}

func (n NullActionHandler) Handle() {
	n.game.Turns[n.game.TurnNumber].Actions = append(n.game.Turns[n.game.TurnNumber].Actions, *n.action)
}

func (NullActionHandler) IsLegit() (bool, error) {
	return true, nil
}
