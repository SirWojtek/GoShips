package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/SirWojtek/GoShips/utilities"
	"github.com/rthornton128/goncurses"
)

var colorMap = map[objects.Color]int16{
	objects.Black:   goncurses.C_BLACK,
	objects.Blue:    goncurses.C_BLUE,
	objects.Cyan:    goncurses.C_CYAN,
	objects.Green:   goncurses.C_GREEN,
	objects.Magenta: goncurses.C_MAGENTA,
	objects.Red:     goncurses.C_RED,
	objects.White:   goncurses.C_WHITE,
	objects.Yellow:  goncurses.C_YELLOW,
}

type objData struct {
	window     *goncurses.Window
	colorIndex int16
}

type gameView struct {
	scene             objects.ObjectInterface
	scaleX, scaleY    float32
	objectToWindowMap map[string]objData
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
		objectToWindowMap: map[string]objData{},
	}
}

func (view *gameView) paint(stdscr *goncurses.Window) {
	paintObjects := view.scene.GetChildsRecursive()

	if len(paintObjects) < len(view.objectToWindowMap) {
		view.removeUnusedWindows(paintObjects)
	}

	for _, obj := range paintObjects {
		view.paintObject(obj, stdscr)
	}
}

func (view *gameView) removeUnusedWindows(objList []objects.ObjectInterface) {
	for key, _ := range view.objectToWindowMap {
		found := false
		for _, obj := range objList {
			if obj.GetName() == key {
				found = true
				break
			}
		}

		if !found {
			delete(view.objectToWindowMap, key)
		}
	}
}

func (view *gameView) paintObject(obj objects.ObjectInterface, stdscr *goncurses.Window) {
	y, x := view.convertToScreenCoords(obj.GetRect())

	objData := view.getOrCreateObjectWindow(obj, stdscr)
	objData.window.Erase()
	objData.window.MoveWindow(y, x)
	objData.window.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
	objData.window.NoutRefresh()
}

func (view *gameView) convertToScreenCoords(rect objects.Rect) (int, int) {
	return int(rect.Y * view.scaleY), int(rect.X * view.scaleX)
}

func (view *gameView) getOrCreateObjectWindow(obj objects.ObjectInterface, stdscr *goncurses.Window) objData {
	data, exist := view.objectToWindowMap[obj.GetName()]
	rect := obj.GetRect()

	if !exist {
		win := stdscr.Sub(
			int(rect.Height), int(rect.Width), 0, 0)

		colorIndex := int16(len(view.objectToWindowMap) + 1)
		goncurses.InitPair(colorIndex, goncurses.C_WHITE, colorMap[obj.GetColor()])
		win.SetBackground(goncurses.ColorPair(colorIndex))

		data = objData{win, colorIndex}
		view.objectToWindowMap[obj.GetName()] = data
	}

	return data
}

func (view *gameView) clean() {
	for _, data := range view.objectToWindowMap {
		data.window.Delete()
	}
}
