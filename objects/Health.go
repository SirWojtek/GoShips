package objects

type Health int

func (hp *Health) GetDamage(dmg Health) {
	*hp -= dmg
}
