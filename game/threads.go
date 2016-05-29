package game

import (
	"github.com/SirWojtek/GoShips/controller"
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
	prePaintControllers []controller.Controller,
	vc view.ViewContext,
	postPaintControllers []controller.Controller,
	waitGroup *sync.WaitGroup,
	exitChannel chan bool) {

	defer vc.End()
	defer waitGroup.Done()

	for {
		select {
		case exit := <-exitChannel:
			utilities.Log.Println("View loop exited", exit)
			return
		default:
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
}

func (t *Threads) controllLoop(
	controller controller.Controller,
	waitGroup *sync.WaitGroup,
	exitChannel chan bool) {

	defer waitGroup.Done()
	for {
		select {
		case exit := <-exitChannel:
			utilities.Log.Println("Control loop exited", exit)
			return
		default:
			controller.Tick()
			t.painted.L.Lock()
			t.painted.Wait()
			t.painted.L.Unlock()
		}
	}
}
