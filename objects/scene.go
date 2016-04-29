package objects

type Scene struct {
	*Object
	Ships []*Ship
}

const sceneWidth = 100
const sceneHeight = 100

func NewScene() Scene {
	sceneBounds := Rect{0, 0, sceneHeight, sceneWidth}
	leftShip := NewShip("LeftShip", Rect{0, 30, 10, 10}, sceneBounds)
	rightShip := NewShip("RightShip", Rect{90, 30, 10, 10}, sceneBounds)

	scene := Scene{
		Object: NewObject("Scene", sceneBounds, sceneBounds),
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
