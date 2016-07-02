package objects

import "image/color"

type Scene struct {
	*Object
	Ships []*Ship
}

const sceneWidth = 100
const sceneHeight = 100

var (
	Red  = color.Gray16{0xf00f}
	Blue = color.Gray16{0x00ff}
)

func NewScene() Scene {
	sceneBounds := Rect{0, 0, sceneHeight, sceneWidth}
	leftShip := NewShip("LeftShip", Rect{0, 30, 5, 5}, Red, sceneBounds, true)
	rightShip := NewShip("RightShip", Rect{90, 30, 5, 5}, Blue, sceneBounds, false)

	scene := Scene{
		Object: NewObject("Scene", sceneBounds, color.Black, sceneBounds),
		Ships:  []*Ship{leftShip, rightShip},
	}

	scene.AddChild(leftShip)
	scene.AddChild(rightShip)

	return scene
}

func (obj *Scene) CanMove(x, y float32) bool {
	return false
}

func (obj *Scene) MoveBy(x, y float32) {
	panic("Scene can not be moved")
}
