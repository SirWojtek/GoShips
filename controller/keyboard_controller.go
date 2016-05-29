package controller

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/SirWojtek/GoShips/view"
)

const moveStep = 5.0

type KeyboardController struct {
	keyboard    view.Keyboard
	ship        *objects.Ship
	exitChannel chan bool
}

func NewKeyboardController(
	keyboard view.Keyboard,
	ship *objects.Ship,
	exitChannel chan bool) *KeyboardController {
	return &KeyboardController{
		keyboard:    keyboard,
		ship:        ship,
		exitChannel: exitChannel,
	}
}

func (controller *KeyboardController) Tick() {
	utilities.Log.Printf("Pressed key: %s", controller.keyboard.GetChar())
	switch controller.keyboard.GetChar() {
	case view.Left:
		controller.moveShip(-moveStep, 0)
	case view.Right:
		controller.moveShip(moveStep, 0)
	case view.Up:
		controller.moveShip(0, -moveStep)
	case view.Down:
		controller.moveShip(0, moveStep)
	case view.Shoot:
		controller.ship.Shoot()
	case view.Quit:
		close(controller.exitChannel)
	default:
		break
	}
}

func (controller *KeyboardController) moveShip(x, y float32) {
	if controller.ship.CanMove(x, y) {
		controller.ship.MoveBy(x, y)
	}
}
