package objects

type Ship struct {
	pos Rect
	hp  Health
}

const shipHealth = 100

func NewShip(position Rect) Ship {
	return Ship{
		pos: position,
		hp:  shipHealth,
	}
}

func (ship *Ship) MoveBy(x, y float32) {
	ship.pos.MoveBy(x, y)
}

func (ship *Ship) GetDamage(dmg Health) {
	ship.hp.GetDamage(dmg)
}
