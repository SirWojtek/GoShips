package objects

type GameState struct {
	LeftShip, RightShip Ship
}

func (state *GameState) Paint() {
	state.LeftShip.Paint()
	state.RightShip.Paint()
}

func NewGameState() GameState {
	return GameState{
		LeftShip:  NewShip("LeftShip", Rect{0, 30, 10, 10}),
		RightShip: NewShip("RightShip", Rect{100, 30, 10, 10}),
	}
}
