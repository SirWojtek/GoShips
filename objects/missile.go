package objects

import (
	"engo.io/ecs"
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
const MissileMovementStep = 7

func NewMissile(x, y float32, sceneBounds Rect, engoWorld *ecs.World) *Missile {
	number++
	return &Missile{
		Object: NewObject(fmt.Sprintf("%s%d", missilePrefix, number), Rect{x, y, 5, 5}, color.White, sceneBounds, engoWorld),
	}
}

func IsObjectMissile(obj ObjectInterface) bool {
	return strings.HasPrefix(obj.GetName(), missilePrefix)
}
