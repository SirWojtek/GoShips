package objects

type Ship struct {
	*Object
	Health
}

const shipHealth = 100

func NewShip(name string, position Rect) *Ship {
	return &Ship{
		Object: NewObject(name, position),
		Health: shipHealth,
	}
}
