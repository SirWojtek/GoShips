package game

import (
	"github.com/SirWojtek/GoShips/controller"
	"github.com/SirWojtek/GoShips/objects"
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

func (t *Threads) paintLoop(scene objects.ObjectInterface, waitGroup *sync.WaitGroup) {
	for {
		scene.Paint()
		time.Sleep(1 / FPS * time.Second)
		t.painted.Broadcast()
	}

	waitGroup.Done()
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
