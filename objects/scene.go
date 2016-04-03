package objects

type Scene struct {
	*Object
	ships []*Ship
}

func NewScene() Scene {
	leftShip := NewShip("LeftShip", Rect{0, 30, 10, 10})
	rightShip := NewShip("RightShip", Rect{100, 30, 10, 10})

	scene := Scene{
		Object: NewObject("Scene", Rect{0, 0, 0, 0}),
		ships:  []*Ship{leftShip, rightShip},
	}

	scene.AddChild(leftShip)
	scene.AddChild(rightShip)

	return scene
}

func (obj *Scene) MoveBy(x, y float32) {
	panic("Scene can not be moved")
}
