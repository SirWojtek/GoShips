package view

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

func (key Key) String() string {
	switch key {
	case None:
		return "None"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Shoot:
		return "Shoot"
	case Quit:
		return "Quit"
	default:
		panic("Unrecognized key")
	}
}

var keyMap = map[goncurses.Key]Key{
	0:                   None,
	goncurses.KEY_LEFT:  Left,
	goncurses.KEY_RIGHT: Right,
	goncurses.KEY_UP:    Up,
	goncurses.KEY_DOWN:  Down,
	' ':                 Shoot,
	'q':                 Quit,
}

type Keyboard struct {
	stdscr *goncurses.Window
}

func newKeyboard(stdscr *goncurses.Window) Keyboard {
	stdscr.Keypad(true)
	stdscr.Timeout(0)
	return Keyboard{stdscr}
}

func (keyboard *Keyboard) GetChar() Key {
	return keyMap[keyboard.stdscr.GetChar()]
}
