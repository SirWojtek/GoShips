package objects

import (
	"fmt"
	"github.com/SirWojtek/GoShips/utilities"
	"sync"
)

type Rect struct {
	X, Y          float32
	Width, Height float32
}

type Color int8

const (
	Black Color = iota
	Blue
	Cyan
	Green
	Magenta
	Red
	White
	Yellow
)

type ObjectInterface interface {
	CanMove(x, y float32) bool
	MoveBy(x, y float32)

	AddChild(ObjectInterface)
	GetChilds() []ObjectInterface
	GetChildsRecursive() []ObjectInterface

	GetRect() Rect
	GetName() string
	GetColor() Color

	CollisionCallback(ObjectInterface) bool
}

type Object struct {
	Rect
	Color
	sync.RWMutex
	name        string
	childs      []ObjectInterface
	sceneBounds Rect
}

func NewObject(name string, r Rect, c Color, bounds Rect) *Object {
	return &Object{
		Rect:        r,
		Color:       c,
		name:        name,
		childs:      []ObjectInterface{},
		sceneBounds: bounds,
	}
}

func (obj *Object) CanMove(x, y float32) bool {
	obj.RLock()
	defer obj.RUnlock()

	newX := obj.X + x
	newMaxX := newX + obj.Width
	newY := obj.Y + y
	newMaxY := newY + obj.Height
	return newMaxX <= obj.sceneBounds.X+obj.sceneBounds.Width &&
		newX >= obj.sceneBounds.X &&
		newMaxY <= obj.sceneBounds.Y+obj.sceneBounds.Height &&
		newY >= obj.sceneBounds.Y
}

func (obj *Object) MoveBy(x, y float32) {
	if !obj.CanMove(x, y) {
		panic(obj.name + " goes out of bounds")
	}
	obj.Lock()
	obj.X += x
	obj.Y += y
	obj.Unlock()

	utilities.Log.Println(obj)
}

func (obj *Object) AddChild(o ObjectInterface) {
	obj.Lock()
	defer obj.Unlock()

	obj.childs = append(obj.childs, o)
}

func (obj *Object) GetChilds() []ObjectInterface {
	obj.RLock()
	defer obj.RUnlock()

	return obj.childs
}

func (obj *Object) GetChildsRecursive() []ObjectInterface {
	obj.RLock()
	defer obj.RUnlock()

	result := obj.childs
	for _, child := range obj.childs {
		result = append(result, child.GetChildsRecursive()...)
	}
	return result
}

func (obj *Object) GetRect() Rect {
	obj.RLock()
	defer obj.RUnlock()

	return obj.Rect
}

func (obj *Object) GetName() string {
	obj.RLock()
	defer obj.RUnlock()

	return obj.name
}

func (obj *Object) GetColor() Color {
	obj.RLock()
	defer obj.RUnlock()

	return obj.Color
}

func (obj *Object) CollisionCallback(other ObjectInterface) bool {
	obj.RLock()
	defer obj.RUnlock()

	utilities.Log.Println("Collision:\n%v\n%v", obj, other)
	return true
}

func (obj *Object) String() string {
	obj.RLock()
	defer obj.RUnlock()

	return fmt.Sprintf("%s: %+v", obj.name, obj.Rect)
}

type ByName []ObjectInterface

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
