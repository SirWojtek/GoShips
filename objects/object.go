package objects

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	//"github.com/SirWojtek/GoShips/utilities"
	"image/color"
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
	DeleteChild(ObjectInterface)
	Delete()

	GetParent() ObjectInterface
	SetParent(ObjectInterface)

	GetBasicEntity() *ecs.BasicEntity
	GetSpaceComponent() *common.SpaceComponent
	GetRenderComponent() *common.RenderComponent
	GetName() string

	CollisionCallback(ObjectInterface) bool
}

type Object struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	name        string
	childs      []ObjectInterface
	parent      ObjectInterface
	sceneBounds Rect
	engoWorld   *ecs.World
}

func NewObject(name string, r Rect, c color.Gray16, bounds Rect, engoWorld *ecs.World) *Object {
	return &Object{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{r.X, r.Y},
			Width:    r.Width,
			Height:   r.Height,
		},
		RenderComponent: common.RenderComponent{
			Drawable: common.Rectangle{},
			Color:    c,
		},
		name:        name,
		childs:      []ObjectInterface{},
		sceneBounds: bounds,
		engoWorld:   engoWorld,
	}
}

func (obj *Object) CanMove(x, y float32) bool {
	newX := obj.SpaceComponent.Position.X + x
	newMaxX := newX + obj.SpaceComponent.Width
	newY := obj.SpaceComponent.Position.Y + y
	newMaxY := newY + obj.SpaceComponent.Height
	return newMaxX <= obj.sceneBounds.X+obj.sceneBounds.Width &&
		newX >= obj.sceneBounds.X &&
		newMaxY <= obj.sceneBounds.Y+obj.sceneBounds.Height &&
		newY >= obj.sceneBounds.Y
}

func (obj *Object) MoveBy(x, y float32) {
	if !obj.CanMove(x, y) {
		panic(obj.name + " goes out of bounds")
	}
	obj.SpaceComponent.Position.X += x
	obj.SpaceComponent.Position.Y += y
}

// TODO: function should only takes necessary args, not whole interface
func (obj *Object) AddChild(o ObjectInterface) {
	o.SetParent(obj)
	obj.childs = append(obj.childs, o)
	obj.addToWorld(o)
}

func (obj *Object) addToWorld(o ObjectInterface) {
	for _, system := range obj.engoWorld.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(o.GetBasicEntity(), o.GetRenderComponent(), o.GetSpaceComponent())
		}
	}
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

func (obj *Object) DeleteChild(childToDelete ObjectInterface) {
	for i, child := range obj.childs {
		if child.GetName() == childToDelete.GetName() {
			obj.engoWorld.RemoveEntity(*child.GetBasicEntity())
			obj.childs = append(obj.childs[:i], obj.childs[i+1:]...)
		}
	}
}

func (obj *Object) Delete() {
	obj.GetParent().DeleteChild(obj)
}

func (obj *Object) GetParent() ObjectInterface {
	return obj.parent
}

func (obj *Object) SetParent(p ObjectInterface) {
	obj.parent = p
}

func (obj *Object) GetBasicEntity() *ecs.BasicEntity {
	return &obj.BasicEntity
}

func (obj *Object) GetSpaceComponent() *common.SpaceComponent {
	return &obj.SpaceComponent
}

func (obj *Object) GetRenderComponent() *common.RenderComponent {
	return &obj.RenderComponent
}

func (obj *Object) GetName() string {
	return obj.name
}

func (obj *Object) CollisionCallback(other ObjectInterface) bool {
	return true
}

func (obj *Object) String() string {
	return fmt.Sprintf("%s: %+v", obj.name, obj.RenderComponent)
}

type ByName []ObjectInterface

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
