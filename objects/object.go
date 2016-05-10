package objects

import (
	"fmt"
	"github.com/SirWojtek/GoShips/utilities"
)

type Rect struct {
	X, Y          float32
	Width, Height float32
}

type ObjectInterface interface {
	CanMove(x, y float32) bool
	MoveBy(x, y float32)

	AddChild(ObjectInterface)
	GetChilds() []ObjectInterface
	GetChildsRecursive() []ObjectInterface

	GetRect() Rect
	GetName() string

	CollisionCallback(ObjectInterface) bool
}

type Object struct {
	name string
	Rect
	childs      []ObjectInterface
	sceneBounds Rect
}

func NewObject(name string, r Rect, bounds Rect) *Object {
	return &Object{
		name:        name,
		Rect:        r,
		childs:      []ObjectInterface{},
		sceneBounds: bounds,
	}
}

func (obj *Object) CanMove(x, y float32) bool {
	newX := obj.X + x
	newY := obj.Y + y
	return newX <= obj.sceneBounds.Width &&
		newX >= obj.sceneBounds.X &&
		newY <= obj.sceneBounds.Height &&
		newY >= obj.sceneBounds.Y
}

func (obj *Object) MoveBy(x, y float32) {
	if !obj.CanMove(x, y) {
		panic(obj.name + " goes out of bounds")
	}

	obj.X += x
	obj.Y += y
}

func (obj *Object) AddChild(o ObjectInterface) {
	obj.childs = append(obj.childs, o)
}

func (obj *Object) GetChilds() []ObjectInterface {
	return obj.childs
}

func (obj *Object) GetChildsRecursive() []ObjectInterface {
	result := obj.childs
	for _, child := range obj.childs {
		result = append(result, child.GetChildsRecursive()...)
	}
	return result
}

func (obj *Object) GetRect() Rect {
	return obj.Rect
}

func (obj *Object) GetName() string {
	return obj.name
}

func (obj *Object) CollisionCallback(other ObjectInterface) bool {
	utilities.Log.Println("Collision:\n%v\n%v", obj, other)
	return true
}

func (obj *Object) String() string {
	return fmt.Sprintf("%s: %+v", obj.name, obj.Rect)
}

type ByName []ObjectInterface

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
