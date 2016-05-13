package objects

type Ship struct {
	*Object
	Health
}

const shipHealth = 100

func NewShip(name string, position Rect, color Color, sceneBounds Rect) *Ship {
	return &Ship{
		Object: NewObject(name, position, color, sceneBounds),
		Health: shipHealth,
	}
}

func (ship *Ship) Shoot() {
	ship.AddChild(NewMissile(
		ship.Rect.X+ship.Rect.Width/2,
		ship.Rect.Y+ship.Rect.Height/2,
		ship.sceneBounds))
}
