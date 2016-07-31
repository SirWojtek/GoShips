package system

import (
	"engo.io/ecs"
	"github.com/SirWojtek/GoShips/objects"
	"math/rand"
	"time"
)

type RandomSystem struct {
	ship       *objects.Ship
	random     *rand.Rand
	sleepDelta float32
}

const sleepPeriod = 0.1 // seconds
const shootModulo = 10

func NewRandomSystem(obj *objects.Ship) *RandomSystem {
	return &RandomSystem{
		ship:       obj,
		random:     rand.New(rand.NewSource(time.Now().UnixNano())),
		sleepDelta: 0.0,
	}
}

func (system *RandomSystem) Remove(ecs.BasicEntity) {}

func (system *RandomSystem) Update(dt float32) {
	system.sleepDelta += dt
	if system.sleepDelta >= sleepPeriod {
		system.sleepDelta = 0.0
		system.PerformAction()
	}
}

func (system *RandomSystem) PerformAction() {
	rand := system.random.Int()
	if rand%shootModulo != 0 {
		x, y := system.generateMoveDelta()
		for ; !system.ship.CanMove(x, y); x, y = system.generateMoveDelta() {
		}
		system.ship.MoveBy(x, y)
	} else {
		system.ship.Shoot()
	}
}

func (system *RandomSystem) generateMoveDelta() (float32, float32) {
	x := 2*system.random.Float32()*objects.ShipMovementStep - objects.ShipMovementStep
	y := 2*system.random.Float32()*objects.ShipMovementStep - objects.ShipMovementStep
	return x, y
}
