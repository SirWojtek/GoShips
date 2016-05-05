package view

import "github.com/rthornton128/goncurses"

type viewInterface interface {
	paint(stdscr *goncurses.Window)
	clean()
}
