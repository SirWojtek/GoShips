package objects

type Ship struct {
	Rect
	Health
}

const shipHealth = 100

func NewShip(position Rect) Ship {
	return Ship{
		Rect:   position,
		Health: shipHealth,
	}
}
