package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/rthornton128/goncurses"
)

type gameView struct {
	scene             objects.ObjectInterface
	scaleX, scaleY    float32
	objectToWindowMap map[string]*goncurses.Window
}

func newGameView(scene objects.ObjectInterface, stdscr *goncurses.Window) *gameView {
	sceneRect := scene.GetRect()
	screenMaxY, screenMaxX := stdscr.MaxYX()

	utilities.Log.Printf("Screen size: %d %d\n", screenMaxX, screenMaxY)
	utilities.Log.Printf("Scene size: %f %f\n", sceneRect.Width, sceneRect.Height)

	return &gameView{
		scene:             scene,
		scaleX:            float32(screenMaxX) / sceneRect.Width,
		scaleY:            float32(screenMaxY) / sceneRect.Height,
		objectToWindowMap: map[string]*goncurses.Window{},
	}
}

func (view *gameView) paint(stdscr *goncurses.Window) {
	for _, obj := range view.scene.GetChildsRecursive() {
		view.paintObject(obj, stdscr)
	}
}

func (view *gameView) paintObject(obj objects.ObjectInterface, stdscr *goncurses.Window) {
	objRect := obj.GetRect()
	y, x := view.convertToScreenCoords(objRect)

	objWindow := view.getOrCreateObjectWindow(obj.GetName(), objRect)
	objWindow.Erase()
	objWindow.NoutRefresh()
	objWindow.MoveWindow(y, x)
	objWindow.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
	objWindow.NoutRefresh()
}

func (view *gameView) convertToScreenCoords(rect objects.Rect) (int, int) {
	return int(rect.Y * view.scaleY), int(rect.X * view.scaleX)
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

func (view *gameView) clean() {
	for _, window := range view.objectToWindowMap {
		window.Delete()
	}
}

func createObjectWindow(rect objects.Rect) *goncurses.Window {
	win, err := goncurses.NewWindow(
		int(rect.Height), int(rect.Width), 0, 0)
	if err != nil {
		panic("Cannot create window")
	}
	return win
}
