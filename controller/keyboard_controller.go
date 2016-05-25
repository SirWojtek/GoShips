package controller

import (
	"github.com/SirWojtek/GoShips/input"
	"github.com/SirWojtek/GoShips/objects"
)

const moveStep = 5.0

type KeyboardController struct {
	keyboard input.Keyboard
	ship     *objects.Ship
}

func (controller *KeyboardController) Tick() {
	switch controller.keyboard.GetChar() {
	case input.Left:
		controller.moveShip(0, moveStep)
	case input.Right:
		controller.moveShip(0, -moveStep)
	case input.Up:
		controller.moveShip(moveStep, 0)
	case input.Down:
		controller.moveShip(-moveStep, 0)
	default:
		break
	}
}

func (controller *KeyboardController) moveShip(x, y float32) {
	if controller.ship.IsTurnedRight {
		move(controller.ship, x, y)
	} else {
		move(controller.ship, -x, y)
	}
}

func move(ship *objects.Ship, x, y float32) {
	if ship.CanMove(x, y) {
		ship.MoveBy(x, y)
	}
}
