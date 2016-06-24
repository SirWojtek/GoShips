package objects

import (
	"github.com/SirWojtek/GoShips/utilities"
	"time"
)

type Ship struct {
	*Object
	Health
	IsTurnedRight bool
	canShoot      bool
	shootTimer    *time.Timer
}

const shipHealth = 100
const ShipMovementStep = 1.2
const shipShootPeriod = 500 // ms

func NewShip(name string, position Rect, color Color, sceneBounds Rect, isTurnedRight bool) *Ship {
	return &Ship{
		Object:        NewObject(name, position, color, sceneBounds),
		Health:        shipHealth,
		IsTurnedRight: isTurnedRight,
		canShoot:      true,
	}
}

func (ship *Ship) Shoot() {
	if ship.canShoot {
		x, y := ship.getMissileCoords()
		ship.AddChild(NewMissile(x, y, ship.sceneBounds))
		ship.canShoot = false
		ship.shootTimer = time.AfterFunc(shipShootPeriod*time.Millisecond, ship.resetShoot)
	}
}

func (ship *Ship) resetShoot() {
	ship.canShoot = true
}

func (ship *Ship) getMissileCoords() (float32, float32) {
	if ship.IsTurnedRight {
		return ship.Rect.X + ship.Rect.Width, ship.Rect.Y + ship.Rect.Height/2
	} else {
		return ship.Rect.X - ship.Rect.Width, ship.Rect.Y + ship.Rect.Height/2
	}
}

func (ship *Ship) CollisionCallback(other ObjectInterface) bool {
	ship.Object.Lock()
	defer ship.Object.Unlock()

	if IsObjectMissile(other) {
		ship.Health.GetDamage(missileDamage)
		other.Delete()
		utilities.Log.Printf("%s hited for %d", ship.name, missileDamage)
	}
	return true
}
