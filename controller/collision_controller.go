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

}
