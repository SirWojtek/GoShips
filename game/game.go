package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
	"sync"
)

type Game struct {
	scene                objects.Scene
	shipControllers      []controller.RandomController
	prePaintControllers  []controller.Controller
	postPaintControllers []controller.Controller
}

func NewGame() Game {
	game := Game{
		scene:                objects.NewScene(),
		shipControllers:      []controller.RandomController{},
		prePaintControllers:  []controller.Controller{},
		postPaintControllers: []controller.Controller{},
	}

	for i := range game.scene.Ships {
		game.shipControllers = append(game.shipControllers,
			controller.NewRandomController(game.scene.Ships[i]))
	}

	game.prePaintControllers = append(game.prePaintControllers,
		controller.NewCollisionController(&game.scene))

	return game
}

func (game *Game) Start() {
	var waitGroup sync.WaitGroup
	threads := NewThreads()

	waitGroup.Add(1)
	go threads.paintLoop(&game.scene,
		game.prePaintControllers,
		game.postPaintControllers,
		&waitGroup)

	for i := range game.shipControllers {
		waitGroup.Add(1)
		go threads.controllLoop(&game.shipControllers[i], &waitGroup)
	}

	waitGroup.Wait()
}
