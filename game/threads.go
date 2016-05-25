package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
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
	vc view.ViewContext,
	postPaintControllers []controller.Controller,
	waitGroup *sync.WaitGroup) {

	defer vc.End()
	defer waitGroup.Done()

	for {
		for _, preController := range prePaintControllers {
			preController.Tick()
		}

		utilities.Log.Println("Before view loop")
		vc.ViewLoop()
		utilities.Log.Println("After view loop")

		for _, postController := range postPaintControllers {
			postController.Tick()
		}

		// TODO: dynamically compute sleep period (keep const FPS)
		t.painted.Broadcast()
		time.Sleep(1 / FPS * time.Second)
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
