package objects

type Ship struct {
	*Object
	Health
	IsTurnedRight bool
}

const shipHealth = 100

func NewShip(name string, position Rect, color Color, sceneBounds Rect, isTurnedRight bool) *Ship {
	return &Ship{
		Object:        NewObject(name, position, color, sceneBounds),
		Health:        shipHealth,
		IsTurnedRight: isTurnedRight,
	}
}

func (ship *Ship) Shoot() {
	x, y := ship.getMissileCoords()
	ship.AddChild(NewMissile(x, y, ship.sceneBounds))
}

func (ship *Ship) getMissileCoords() (float32, float32) {
	if ship.IsTurnedRight {
		return ship.Rect.X + ship.Rect.Width, ship.Rect.Y + ship.Rect.Height/2
	} else {
		return ship.Rect.X - ship.Rect.Width, ship.Rect.Y + ship.Rect.Height/2
	}
}
