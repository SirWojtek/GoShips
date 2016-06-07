package utilities

type BroadcastChannel struct {
	Out       chan bool
	listeners int
}

func NewBroadcastChannel() *BroadcastChannel {
	return &BroadcastChannel{
		Out:       make(chan bool),
		listeners: 0,
	}
}

func (bc *BroadcastChannel) AddListener() {
	bc.listeners += 1
}

func (bc *BroadcastChannel) Broadcast(value bool) {
	for i := 0; i < bc.listeners; i++ {
		bc.Out <- value
	}
}
