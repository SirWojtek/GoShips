package system

import (
	"engo.io/ecs"
	"github.com/SirWojtek/GoShips/objects"
)

type MissileSystem struct {
	leftShip  objects.ObjectInterface
	rightShip objects.ObjectInterface
}

func NewMissileSystem(left, right objects.ObjectInterface) *MissileSystem {
	return &MissileSystem{left, right}
}

func (system *MissileSystem) Remove(ecs.BasicEntity) {}

func (system *MissileSystem) Update(dt float32) {
	moveShipMissiles(system.leftShip, 1)
	moveShipMissiles(system.rightShip, -1)
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
