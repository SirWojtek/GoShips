package view

import "github.com/rthornton128/goncurses"

type ViewContext struct {
	stdscr *goncurses.Window
}

func NewViewContext() ViewContext {
	scr, err := goncurses.Init()
	if err != nil {
		panic("ncurses init failed")
	}
	return ViewContext{
		stdscr: scr,
	}
}

func End() {
	goncurses.End()
}

func (context ViewContext) ViewLoop() {
	context.stdscr.Print("Hello World!")
	context.stdscr.Refresh()
}
