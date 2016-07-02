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

func NewGame() Game {
	utilities.Init()
	game := Game{
		scene: objects.NewScene(),
		//shipControllers:      []controller.RandomController{},
		//prePaintControllers:  []controller.Controller{},
		//postPaintControllers: []controller.Controller{},
		exitChannel: utilities.NewBroadcastChannel(),
		engoOpts: engo.RunOptions{
			Title:  "GoShips",
			Width:  800,
			Height: 600,
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
	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(system.NewCollisionSystem(&game.scene))
}

func (game *Game) Start() {
	engo.Run(game.engoOpts, game)

	//var waitGroup sync.WaitGroup
	//threads := NewThreads()

	//waitGroup.Add(1)
	//go threads.paintLoop(
	//game.prePaintControllers,
	//game.viewContext,
	//game.postPaintControllers,
	//&waitGroup,
	//game.exitChannel)

	//for i := range game.shipControllers {
	//waitGroup.Add(1)
	//go threads.controllLoop(
	//&game.shipControllers[i],
	//&waitGroup,
	//game.exitChannel)
	//}

	//utilities.Log.Println("Threads started")
	//waitGroup.Wait()
	//utilities.Log.Println("Threads stoped")
}
