package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/view"
	"sync"
	"time"
)

type Threads struct {
	mutex   sync.Mutex
	painted *sync.Cond
}

const FPS = 1

func NewThreads() Threads {
	var threads Threads
	threads.painted = sync.NewCond(&threads.mutex)
	return threads
}

func (t *Threads) paintLoop(
	scene objects.ObjectInterface,
	prePaintControllers []controller.Controller,
	postPaintControllers []controller.Controller,
	waitGroup *sync.WaitGroup) {

	vc := view.NewViewContext()
	defer view.End()
	defer waitGroup.Done()

	for {
		for _, preController := range prePaintControllers {
			preController.Tick()
		}

		vc.ViewLoop()

		for _, postController := range postPaintControllers {
			postController.Tick()
		}

		// TODO: dynamically compute sleep period (keep const FPS)
		time.Sleep(1 / FPS * time.Second)
		t.painted.Broadcast()
	}
}

func (t *Threads) controllLoop(controller controller.Controller, waitGroup *sync.WaitGroup) {
	for {
		controller.Tick()
		t.painted.L.Lock()
		t.painted.Wait()
		t.painted.L.Unlock()
	}

	waitGroup.Done()
}
