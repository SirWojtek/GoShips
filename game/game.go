package game

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/system"
	"github.com/SirWojtek/GoShips/utilities"
	"image/color"
)

type Game struct {
	scene objects.Scene
	//shipControllers      []controller.RandomController
	//prePaintControllers  []controller.Controller
	//viewContext          view.ViewContext
	//postPaintControllers []controller.Controller
	exitChannel *utilities.BroadcastChannel
	engoOpts    engo.RunOptions
}

const sceneWidth = 800
const sceneHeight = 600

func NewGame() Game {
	utilities.Init()
	game := Game{
		scene: objects.NewScene(sceneWidth, sceneHeight),
		//shipControllers:      []controller.RandomController{},
		//prePaintControllers:  []controller.Controller{},
		//postPaintControllers: []controller.Controller{},
		exitChannel: utilities.NewBroadcastChannel(),
		engoOpts: engo.RunOptions{
			Title:         "GoShips",
			Width:         sceneWidth,
			Height:        sceneHeight,
			ScaleOnResize: true,
		},
	}
	//game.viewContext = view.NewViewContext(&game.scene)

	//for i := range game.scene.Ships {
	//game.shipControllers = append(game.shipControllers,
	//controller.NewRandomController(game.scene.Ships[i]))
	//}

	//game.shipControllers = append(game.shipControllers,
	//controller.NewRandomController(game.scene.Ships[1]))

	//game.prePaintControllers = append(game.prePaintControllers,
	//controller.NewCollisionController(&game.scene))
	//game.prePaintControllers = append(game.prePaintControllers,
	//controller.NewPreGameController(game.scene.Ships[0], game.scene.Ships[1]))
	//game.prePaintControllers = append(game.prePaintControllers,
	//controller.NewKeyboardController(
	//game.viewContext.Keyboard, game.scene.Ships[0], game.exitChannel))

	return game
}

func (*Game) Type() string { return "GoShips" }
func (*Game) Preload()     {}
func (game *Game) Setup(world *ecs.World) {
	common.SetBackground(color.Black)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(system.NewCollisionSystem(&game.scene))
	world.AddSystem(system.NewRandomSystem(game.scene.Ships[1])) // NOTE: only right ship random
	world.AddSystem(system.NewMissileSystem(game.scene.Ships[0], game.scene.Ships[1]))

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			addSceneToRenderComponent(sys, &game.scene)
		}
	}
}

func addSceneToRenderComponent(sys *common.RenderSystem, scene *objects.Scene) {
	for _, obj := range scene.GetChildsRecursive() {
		sys.Add(obj.GetBasicEntity(), obj.GetRenderComponent(), obj.GetSpaceComponent())
	}
}

func (game *Game) Start() {
	engo.Run(game.engoOpts, game)
}
