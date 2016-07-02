package objects

import (
	//"github.com/SirWojtek/GoShips/utilities"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"image/color"
	"sync"
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
	sync.RWMutex
	name        string
	childs      []ObjectInterface
	parent      ObjectInterface
	sceneBounds Rect
}

func NewObject(name string, r Rect, c color.Gray16, bounds Rect) *Object {
	return &Object{
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
	}
}

func (obj *Object) CanMove(x, y float32) bool {
	obj.RLock()
	defer obj.RUnlock()

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
	obj.Lock()
	obj.SpaceComponent.Position.X += x
	obj.SpaceComponent.Position.Y += y
	obj.Unlock()
}

func (obj *Object) AddChild(o ObjectInterface) {
	obj.Lock()
	defer obj.Unlock()

	o.SetParent(obj)
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

func (obj *Object) DeleteChild(childToDelete ObjectInterface) {
	for i, child := range obj.childs {
		if child.GetName() == childToDelete.GetName() {
			obj.childs = append(obj.childs[:i], obj.childs[i+1:]...)
		}
	}
}

func (obj *Object) Delete() {
	obj.GetParent().DeleteChild(obj)
}

func (obj *Object) GetParent() ObjectInterface {
	obj.RLock()
	defer obj.RUnlock()

	return obj.parent
}

func (obj *Object) SetParent(p ObjectInterface) {
	obj.Lock()
	defer obj.Unlock()

	obj.parent = p
}

func (obj *Object) GetBasicEntity() *ecs.BasicEntity {
	obj.RLock()
	defer obj.RUnlock()

	return &obj.BasicEntity
}

func (obj *Object) GetSpaceComponent() *common.SpaceComponent {
	obj.RLock()
	defer obj.RUnlock()

	return &obj.SpaceComponent
}

func (obj *Object) GetRenderComponent() *common.RenderComponent {
	obj.RLock()
	defer obj.RUnlock()

	return &obj.RenderComponent
}

func (obj *Object) GetName() string {
	obj.RLock()
	defer obj.RUnlock()

	return obj.name
}

func (obj *Object) CollisionCallback(other ObjectInterface) bool {
	obj.RLock()
	defer obj.RUnlock()
	return true
}

func (obj *Object) String() string {
	obj.RLock()
	defer obj.RUnlock()

	return fmt.Sprintf("%s: %+v", obj.name, obj.RenderComponent)
}

type ByName []ObjectInterface

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
