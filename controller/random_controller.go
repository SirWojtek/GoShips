package controller

import (
	"github.com/SirWojtek/GoShips/objects"
	"math/rand"
	"time"
)

type RandomController struct {
	ship   *objects.Ship
	random *rand.Rand
}

const maxX = 10
const maxY = 10

func NewRandomController(obj *objects.Ship) RandomController {
	return RandomController{
		ship:   obj,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (controller *RandomController) generateMoveDelta() (float32, float32) {
	x := 2*controller.random.Float32()*maxX - maxX
	y := 2*controller.random.Float32()*maxY - maxY
	return x, y
}

func (controller *RandomController) Tick() {
	x, y := controller.generateMoveDelta()

	for ; !controller.ship.CanMove(x, y); x, y = controller.generateMoveDelta() {
	}

	if int(x)%3 != 0 {
		controller.ship.MoveBy(x, y)
	} else {
		controller.ship.Shoot()
	}
}
