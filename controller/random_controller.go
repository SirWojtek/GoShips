package controller

import (
	"github.com/SirWojtek/GoShips/objects"
	"math/rand"
	"time"
)

type RandomController struct {
	object objects.ObjectInterface
	random *rand.Rand
}

const maxX = 10
const maxY = 10

func NewRandomController(obj objects.ObjectInterface) RandomController {
	return RandomController{
		object: obj,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (controller *RandomController) Tick() {
	x := controller.random.Float32() * maxX
	y := controller.random.Float32() * maxY
	controller.object.MoveBy(x, y)
}
