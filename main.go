package main

//import "github.com/SirWojtek/GoShips/game"

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

type scene struct{}

func (*scene) Type() string { return "Hello" }
func (*scene) Preload()     {}
func (*scene) Setup(world *ecs.World) {
	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
}

func main() {
	//game := game.NewGame()
	//game.Start()

	opts := engo.RunOptions{
		Title:  "Hello World!",
		Width:  800,
		Height: 600,
	}
	engo.Run(opts, &scene{})
}
