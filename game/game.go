package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/SirWojtek/GoShips/view"
	"sync"
)

type Game struct {
	scene                objects.Scene
	shipControllers      []controller.RandomController
	prePaintControllers  []controller.Controller
	viewContext          view.ViewContext
	postPaintControllers []controller.Controller
	exitChannel          *utilities.BroadcastChannel
}

const logFile = "log.txt"

func NewGame() Game {
	utilities.Init(logFile)
	game := Game{
		scene:                objects.NewScene(),
		shipControllers:      []controller.RandomController{},
		prePaintControllers:  []controller.Controller{},
		postPaintControllers: []controller.Controller{},
		exitChannel:          utilities.NewBroadcastChannel(),
	}
	game.viewContext = view.NewViewContext(&game.scene)

	//for i := range game.scene.Ships {
	//game.shipControllers = append(game.shipControllers,
	//controller.NewRandomController(game.scene.Ships[i]))
	//}

	game.shipControllers = append(game.shipControllers,
		controller.NewRandomController(game.scene.Ships[1]))

	game.prePaintControllers = append(game.prePaintControllers,
		controller.NewCollisionController(&game.scene))
	game.prePaintControllers = append(game.prePaintControllers,
		controller.NewPreGameController(game.scene.Ships[0], game.scene.Ships[1]))
	game.prePaintControllers = append(game.prePaintControllers,
		controller.NewKeyboardController(
			game.viewContext.Keyboard, game.scene.Ships[0], game.exitChannel))

	return game
}

func (game *Game) Start() {
	var waitGroup sync.WaitGroup
	threads := NewThreads()

	waitGroup.Add(1)
	go threads.paintLoop(
		game.prePaintControllers,
		game.viewContext,
		game.postPaintControllers,
		&waitGroup,
		game.exitChannel)

	for i := range game.shipControllers {
		waitGroup.Add(1)
		go threads.controllLoop(
			&game.shipControllers[i],
			&waitGroup,
			game.exitChannel)
	}

	utilities.Log.Println("Threads started")
	waitGroup.Wait()
	utilities.Log.Println("Threads stoped")
}
