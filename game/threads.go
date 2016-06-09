package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/SirWojtek/GoShips/view"
	"sync"
	"time"
)

type Threads struct {
	painted chan bool
}

const FPS = 10

func NewThreads() Threads {
	var threads Threads
	threads.painted = make(chan bool)
	return threads
}

func (t *Threads) paintLoop(
	prePaintControllers []controller.Controller,
	vc view.ViewContext,
	postPaintControllers []controller.Controller,
	waitGroup *sync.WaitGroup,
	exitChannel *utilities.BroadcastChannel) {

	exitChannel.AddListener()
	defer vc.End()
	defer waitGroup.Done()
	defer close(t.painted)

	for {
		select {
		case exit := <-exitChannel.Out:
			utilities.Log.Println("View loop exited", exit)
			return
		default:
			for _, preController := range prePaintControllers {
				preController.Tick()
			}

			vc.ViewLoop()

			for _, postController := range postPaintControllers {
				postController.Tick()
			}

			select {
			case t.painted <- true:
			default:
				// TODO: dynamically compute sleep period (keep const FPS)
				time.Sleep(1 / FPS * time.Second)
			}
		}
	}
}

func (t *Threads) controllLoop(
	controller controller.Controller,
	waitGroup *sync.WaitGroup,
	exitChannel *utilities.BroadcastChannel) {

	exitChannel.AddListener()
	defer waitGroup.Done()

	for {
		select {
		case exit := <-exitChannel.Out:
			utilities.Log.Println("Control loop exited", exit)
			return
		case painted := <-t.painted:
			if painted {
				controller.Tick()
			}
		}
	}
}
