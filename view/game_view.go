package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/rthornton128/goncurses"
)

type gameView struct {
	scene             objects.ObjectInterface
	objectToWindowMap map[string]*goncurses.Window
}

func newGameView(scene objects.ObjectInterface) *gameView {
	return &gameView{
		scene:             scene,
		objectToWindowMap: map[string]*goncurses.Window{},
	}
}

func (view *gameView) paint(stdscr *goncurses.Window) {
	for _, obj := range view.scene.GetChildsRecursive() {
		utilities.Log.Println(obj)
		view.paintObject(obj)
	}
}

func (view *gameView) paintObject(obj objects.ObjectInterface) {
	objRect := obj.GetRect()
	objWindow := view.getOrCreateObjectWindow(obj.GetName(), objRect)

	objWindow.Erase()
	objWindow.NoutRefresh()
	objWindow.MoveWindow(int(objRect.Y), int(objRect.X))
	objWindow.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
	objWindow.NoutRefresh()
}

func (view *gameView) getOrCreateObjectWindow(name string, rect objects.Rect) *goncurses.Window {
	window, exist := view.objectToWindowMap[name]

	if !exist {
		win := createObjectWindow(rect)
		utilities.Log.Println("Created window for: " + name)
		view.objectToWindowMap[name] = win
		window = win
	}

	return window
}

func createObjectWindow(rect objects.Rect) *goncurses.Window {
	win, err := goncurses.NewWindow(
		int(rect.Height), int(rect.Width),
		int(rect.Y), int(rect.X))
	if err != nil {
		panic("Cannot create window")
	}
	return win
}

func (view *gameView) clean() {
	for _, window := range view.objectToWindowMap {
		window.Delete()
	}
}
