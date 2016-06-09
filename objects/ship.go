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
	shootTimer    <-chan time.Time
}

const shipHealth = 100
const ShipMovementStep = 1.2
const shipShootPeriod = 500 // ms

func newShootTimer() <-chan time.Time {
	return time.After(shipShootPeriod * time.Millisecond)
}

func NewShip(name string, position Rect, color Color, sceneBounds Rect, isTurnedRight bool) *Ship {
	return &Ship{
		Object:        NewObject(name, position, color, sceneBounds),
		Health:        shipHealth,
		IsTurnedRight: isTurnedRight,
		canShoot:      true,
		shootTimer:    newShootTimer(),
	}
}

func (ship *Ship) Shoot() {
	select {
	case <-ship.shootTimer:
		ship.canShoot = true
	default:
		if ship.canShoot {
			x, y := ship.getMissileCoords()
			ship.AddChild(NewMissile(x, y, ship.sceneBounds))
			ship.canShoot = false
			ship.shootTimer = newShootTimer()
		}
	}
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
		utilities.Log.Printf("Ship hited for %d", missileDamage)
	}
	return true
}
