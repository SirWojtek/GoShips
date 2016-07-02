package system

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"github.com/SirWojtek/GoShips/objects"
)

type CollisionSystem struct {
	scene objects.ObjectInterface
}

func NewCollisionSystem(scene objects.ObjectInterface) *CollisionSystem {
	return &CollisionSystem{
		scene: scene,
	}
}

func (system *CollisionSystem) Remove(ecs.BasicEntity) {}

func (controller *CollisionSystem) Update(dt float32) {
	allObj := controller.scene.GetChildsRecursive()

	for i := 0; i < len(allObj); i++ {
		for j := i; j < len(allObj); j++ {
			if areRectsCollide(allObj[i].GetSpaceComponent(), allObj[j].GetSpaceComponent()) {
				allObj[i].CollisionCallback(allObj[j])
				allObj[j].CollisionCallback(allObj[i])
			}
		}
	}
}

func areRectsCollide(a, b *common.SpaceComponent) bool {
	aMaxX, aMaxY := getMaxRectCoords(a)
	bMaxX, bMaxY := getMaxRectCoords(b)
	return (isBetween(b.Position.X, a.Position.X, aMaxX) && isBetween(b.Position.Y, a.Position.Y, aMaxY)) ||
		(isBetween(bMaxX, a.Position.X, aMaxX) && isBetween(bMaxY, a.Position.Y, aMaxY))
}

func getMaxRectCoords(r *common.SpaceComponent) (float32, float32) {
	rMaxX := r.Position.X + r.Width
	rMaxY := r.Position.Y + r.Height
	return rMaxX, rMaxY
}

func isBetween(x, a, b float32) bool {
	return x > a && x < b
}
