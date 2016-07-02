package system

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/SirWojtek/GoShips/view"
)

type KeyboardController struct {
	keyboard    view.Keyboard
	ship        *objects.Ship
	exitChannel *utilities.BroadcastChannel
}

func NewKeyboardController(
	keyboard view.Keyboard,
	ship *objects.Ship,
	exitChannel *utilities.BroadcastChannel) *KeyboardController {
	return &KeyboardController{
		keyboard:    keyboard,
		ship:        ship,
		exitChannel: exitChannel,
	}
}

func (controller *KeyboardController) Tick() {
	select {
	case key := <-controller.keyboard.KeyChannel:
		switch key {
		case view.Left:
			controller.moveShip(-objects.ShipMovementStep, 0)
		case view.Right:
			controller.moveShip(objects.ShipMovementStep, 0)
		case view.Up:
			controller.moveShip(0, -objects.ShipMovementStep)
		case view.Down:
			controller.moveShip(0, objects.ShipMovementStep)
		case view.Shoot:
			controller.ship.Shoot()
		case view.Quit:
			controller.exitChannel.Broadcast(true)
		default:
			break
		}
	default:
	}
}

func (controller *KeyboardController) moveShip(x, y float32) {
	if controller.ship.CanMove(x, y) {
		controller.ship.MoveBy(x, y)
	}
}
