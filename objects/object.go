package objects

import "fmt"

type Rect struct {
	X, Y          float32
	Width, Height float32
}

type ObjectInterface interface {
	CanMove(x, y float32) bool
	MoveBy(x, y float32)
	Paint()
	AddChild(ObjectInterface)
	GetChilds() []ObjectInterface
	GetChildsRecursive() []ObjectInterface
	GetRect() Rect
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

func (obj *Object) Paint() {
	// TODO: implement paint
	fmt.Println(obj)
	for _, child := range obj.childs {
		child.Paint()
	}
}

func (obj *Object) AddChild(o ObjectInterface) {
	obj.childs = append(obj.childs, o)
}

func (obj *Object) GetChilds() []ObjectInterface {
	return obj.childs
}

func (obj *Object) GetChildsRecursive() []ObjectInterface {
	result := []ObjectInterface{}
	for _, child := range obj.childs {
		result = append(result, child.GetChilds()...)
	}
	return result
}

func (obj *Object) GetRect() Rect {
	return obj.Rect
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

func (obj *Object) CollisionCallback(other *Object) bool {
	fmt.Println("Collision:\n%v\n%v", obj, other)
	return true
}

func (obj *Object) String() string {
	return fmt.Sprintf("%s: %+v", obj.name, obj.Rect)
}

type Health int

func (hp *Health) GetDamage(dmg Health) {
	*hp -= dmg
}
