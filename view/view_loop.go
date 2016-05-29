package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/rthornton128/goncurses"
)

type ViewContext struct {
	stdscr      *goncurses.Window
	currentView viewInterface
	Keyboard    Keyboard
}

func NewViewContext(sc objects.ObjectInterface) ViewContext {
	scr, err := goncurses.Init()
	if err != nil {
		panic("ncurses init failed")
	}

	goncurses.Echo(false)
	goncurses.CBreak(true)
	goncurses.Cursor(0)
	if err := goncurses.StartColor(); err != nil {
		panic("Color can not be enabled")
	}

	return ViewContext{
		stdscr:      scr,
		currentView: newGameView(sc, scr),
		Keyboard:    newKeyboard(scr),
	}
}

func (context ViewContext) End() {
	context.currentView.clean()
	goncurses.End()
}

func (context ViewContext) ViewLoop() {
	context.stdscr.Erase()
	context.stdscr.NoutRefresh()
	context.currentView.paint(context.stdscr)
	goncurses.Update()
}
