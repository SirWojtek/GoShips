package objects

type Rect struct {
	X, Y          float32
	Width, Height float32
}
type Health int

type Moveable interface {
	MoveBy(x, y float32)
}
type Damageable interface {
	GetDamage(Health)
}

func (r *Rect) MoveBy(x, y float32) {
	r.X += x
	r.Y += y
}

func (hp *Health) GetDamage(dmg Health) {
	*hp -= dmg
}
