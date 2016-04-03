package objects

type Scene struct {
	*Object
}

func NewScene() Scene {
	scene := Scene{
		Object: NewObject("Scene", Rect{0, 0, 0, 0}),
	}

	scene.AddChild(NewShip("LeftShip", Rect{0, 30, 10, 10}))
	scene.AddChild(NewShip("RightShip", Rect{100, 30, 10, 10}))

	return scene
}
