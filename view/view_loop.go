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
	return ViewContext{
		stdscr: scr,
		currentView: &gameView{
			scene: sc,
		},
	}
}

func End() {
	goncurses.End()
}

func (context ViewContext) ViewLoop() {
	context.currentView.paint(context.stdscr)
	context.stdscr.Refresh()
}
