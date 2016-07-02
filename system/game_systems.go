package system

import (
	"github.com/SirWojtek/GoShips/objects"
)

type PreGameController struct {
	leftShip  objects.ObjectInterface
	rightShip objects.ObjectInterface
}

func NewPreGameController(left, right objects.ObjectInterface) *PreGameController {
	return &PreGameController{left, right}
}

func (controller *PreGameController) Tick() {
	controller.moveMissiles()
}

func (controller *PreGameController) moveMissiles() {
	moveShipMissiles(controller.leftShip, 1)
	moveShipMissiles(controller.rightShip, -1)
}

func moveShipMissiles(ship objects.ObjectInterface, sign float32) {
	for _, shipChild := range ship.GetChilds() {
		if objects.IsObjectMissile(shipChild) {
			handleMissile(shipChild, sign)
		}
	}
}

func handleMissile(missile objects.ObjectInterface, sign float32) {
	if missile.CanMove(objects.MissileMovementStep*sign, 0) {
		missile.MoveBy(objects.MissileMovementStep*sign, 0)
	} else {
		missile.Delete()
	}
}
