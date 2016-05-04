package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/rthornton128/goncurses"
)

type gameView struct {
	scene objects.ObjectInterface
}

func (view *gameView) paint(stdscr *goncurses.Window) {
	stdscr.Print("Hello World!")
}
