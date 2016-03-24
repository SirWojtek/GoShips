package objects

import "fmt"

type Rect struct {
	X, Y          float32
	Width, Height float32
}

type PaintComponent interface {
	Paint()
	AddChild(PaintComponent)
}

type Object struct {
	name string
	Rect
	childs []PaintComponent
}

func NewObject(name string, r Rect) Object {
	return Object{
		name:   name,
		Rect:   r,
		childs: []PaintComponent{},
	}
}

func (obj *Object) Paint() {
	// TODO: implement paint
	fmt.Println(obj)
	for _, child := range obj.childs {
		child.Paint()
	}
}

func (obj *Object) AddChild(c PaintComponent) {
	obj.childs = append(obj.childs, c)
}

func (obj *Object) MoveBy(x, y float32) {
	obj.X += x
	obj.Y += y
}

func (obj *Object) String() string {
	return fmt.Sprintf("%s: %+v", obj.name, obj.Rect)
}

type Health int

func (hp *Health) GetDamage(dmg Health) {
	*hp -= dmg
}
