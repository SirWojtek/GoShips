package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
	"sync"
	"time"
)

const FPS = 1 / 60

func paintLoop(scene objects.ObjectInterface, waitGroup *sync.WaitGroup) {
	for {
		scene.Paint()
		time.Sleep(FPS * time.Second)
	}

	waitGroup.Done()
}

func controllLoop(controller controller.Controller, waitGroup *sync.WaitGroup) {
	for {
		controller.Tick()
	}

	waitGroup.Done()
}
