package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
	"sync"
)

type Game struct {
	scene           objects.Scene
	shipControllers []controller.RandomController
}

func NewGame() Game {
	game := Game{
		scene:           objects.NewScene(),
		shipControllers: []controller.RandomController{},
	}

	for i := range game.scene.Ships {
		game.shipControllers = append(game.shipControllers,
			controller.NewRandomController(game.scene.Ships[i]))
	}

	return game
}

func (game *Game) Start() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go paintLoop(&game.scene, &waitGroup)

	for i := range game.shipControllers {
		waitGroup.Add(1)
		go controllLoop(&game.shipControllers[i], &waitGroup)
	}

	waitGroup.Wait()
}
