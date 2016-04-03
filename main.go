package main

import "github.com/SirWojtek/GoShips/objects"
import "github.com/SirWojtek/GoShips/controller"

func main() {
	scene := objects.NewScene()
	controller := controller.NewRandomController(scene.Ships[0])

	scene.Paint()
	controller.Tick()
	scene.Paint()
}
