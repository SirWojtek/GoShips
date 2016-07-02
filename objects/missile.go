package objects

import (
	"fmt"
	"image/color"
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
		Object: NewObject(fmt.Sprintf("%s%d", missilePrefix, number), Rect{x, y, 1, 1}, color.White, sceneBounds),
	}
}

func IsObjectMissile(obj ObjectInterface) bool {
	return strings.HasPrefix(obj.GetName(), missilePrefix)
}
