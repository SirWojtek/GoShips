package input

import (
	"github.com/rthornton128/goncurses"
)

type Key int8

const (
	None Key = iota
	Left
	Right
	Up
	Down
	Shoot
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

type Keyboard struct {
	stdscr *goncurses.Window
}

func NewKeyboard(stdscr *goncurses.Window) Keyboard {
	goncurses.CBreak(false)
	goncurses.Echo(false)
	stdscr.Keypad(true)
	stdscr.Timeout(0)
	return Keyboard{stdscr}
}

func (keyboard *Keyboard) GetChar() Key {
	return keyMap[keyboard.stdscr.GetChar()]
}
