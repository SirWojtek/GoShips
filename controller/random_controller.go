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
	x := 2*controller.random.Float32()*maxX - maxX
	y := 2*controller.random.Float32()*maxY - maxY
	controller.object.MoveBy(x, y)
}
