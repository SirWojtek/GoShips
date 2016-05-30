package controller

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
)

type PreGameController struct {
	leftShip  objects.ObjectInterface
	rightShip objects.ObjectInterface
}

const missileVelocity = .0015

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
	if missile.CanMove(missileVelocity*sign, 0) {
		missile.MoveBy(missileVelocity*sign, 0)
	} else {
		utilities.Log.Println("Remove missile: " + missile.GetName())
		missile.Delete()
	}
}
