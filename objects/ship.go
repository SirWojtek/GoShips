package objects

type Ship struct {
	*Object
	Health
}

const shipHealth = 100

func NewShip(name string, position Rect, sceneBounds Rect) *Ship {
	return &Ship{
		Object: NewObject(name, position, sceneBounds),
		Health: shipHealth,
	}
}
