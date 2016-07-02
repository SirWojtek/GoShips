package controller

import (
	"github.com/SirWojtek/GoShips/objects"
)

type CollisionController struct {
	scene objects.ObjectInterface
}

func NewCollisionController(scene objects.ObjectInterface) *CollisionController {
	return &CollisionController{
		scene: scene,
	}
}

func (controller *CollisionController) Tick() {
	allObj := controller.scene.GetChildsRecursive()

	for i := 0; i < len(allObj); i++ {
		for j := i; j < len(allObj); j++ {
			if areRectsCollide(allObj[i].GetRect(), allObj[j].GetRect()) {
				allObj[i].CollisionCallback(allObj[j])
				allObj[j].CollisionCallback(allObj[i])
			}
		}
	}
}

func areRectsCollide(a, b objects.Rect) bool {
	aMaxX, aMaxY := getMaxRectCoords(a)
	bMaxX, bMaxY := getMaxRectCoords(b)
	return (isBetween(b.X, a.X, aMaxX) && isBetween(b.Y, a.Y, aMaxY)) ||
		(isBetween(bMaxX, a.X, aMaxX) && isBetween(bMaxY, a.Y, aMaxY))
}

func getMaxRectCoords(r objects.Rect) (float32, float32) {
	rMaxX := r.X + r.Width
	rMaxY := r.Y + r.Height
	return rMaxX, rMaxY
}

func isBetween(x, a, b float32) bool {
	return x > a && x < b
}
