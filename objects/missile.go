package objects

import "strings"

type Missile struct {
	*Object
}

var number = 0

const missilePrefix = "missile"

func NewMissile(x, y float32, sceneBounds Rect) *Missile {
	number++
	return &Missile{
		Object: NewObject(missilePrefix+string(number), Rect{x, y, 1, 1}, White, sceneBounds),
	}
}

func IsObjectMissile(obj ObjectInterface) bool {
	return strings.HasPrefix(obj.GetName(), missilePrefix)
}
