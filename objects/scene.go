package objects

import "image/color"

type Scene struct {
	*Object
	Ships []*Ship
}

var (
	Red  = color.Gray16{0xff0f}
	Blue = color.Gray16{0xf00f}
)

func NewScene(sceneWidth, sceneHeight float32) Scene {
	sceneBounds := Rect{0, 0, sceneWidth, sceneHeight}
	leftShip := NewShip("LeftShip", Rect{5, sceneHeight / 2, 15, 20}, Red, sceneBounds, true)
	rightShip := NewShip("RightShip", Rect{sceneWidth - 20, sceneHeight / 2, 15, 20}, Blue, sceneBounds, false)

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
