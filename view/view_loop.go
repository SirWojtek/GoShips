package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/rthornton128/goncurses"
)

type ViewContext struct {
	stdscr      *goncurses.Window
	currentView viewInterface
}

func NewViewContext(sc objects.ObjectInterface) ViewContext {
	scr, err := goncurses.Init()
	if err != nil {
		panic("ncurses init failed")
	}

	goncurses.Echo(false)
	goncurses.CBreak(true)
	goncurses.Cursor(0)

	return ViewContext{
		stdscr:      scr,
		currentView: newGameView(sc),
	}
}

func (context ViewContext) End() {
	context.currentView.clean()
	goncurses.End()
}

func (context ViewContext) ViewLoop() {
	context.currentView.paint(context.stdscr)
	goncurses.Update()
	//context.stdscr.Refresh()
}
