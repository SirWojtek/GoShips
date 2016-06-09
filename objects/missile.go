package objects

import (
	"fmt"
	"strings"
)

type Missile struct {
	*Object
}

var number = 0

const missilePrefix = "missile"
const missileDamage = 20
const MissileMovementStep = .005

func NewMissile(x, y float32, sceneBounds Rect) *Missile {
	number++
	return &Missile{
		Object: NewObject(fmt.Sprintf("%s%d", missilePrefix, number), Rect{x, y, 1, 1}, White, sceneBounds),
	}
}

func IsObjectMissile(obj ObjectInterface) bool {
	return strings.HasPrefix(obj.GetName(), missilePrefix)
}
