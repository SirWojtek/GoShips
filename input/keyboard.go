package input

import (
	"github.com/rthornton128/goncurses"
)

type Key int8

const (
	Key None = iota
	Left
	Right
	Up
	Down
	Shot
	Quit
)

var keyMap = map[goncurses.Key]Key{
	0:                   None,
	goncurses.KEY_LEFT:  Left,
	goncurses.KEY_RIGHT: Right,
	goncurses.KEY_UP:    Up,
	goncurses.KEY_DOWN:  Down,
	goncurses.KEY_ENTER: Shoot,
	goncurses.KEY_EXIT:  Quit,
}

func InitKeyboard(stdscr *goncurses.Window) {
	goncurses.CBreak(false)
	goncurses.Echo(false)
	stdscr.Keypad(true)
	stdscr.Timeout(0)
}

func GetChar(stdscr *goncurses.Window) Key {
	return keyMap[stdscr.GetChar()]
}
