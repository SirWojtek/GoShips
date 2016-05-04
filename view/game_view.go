package view

import (
	"github.com/SirWojtek/GoShips/objects"
	"github.com/rthornton128/goncurses"
)

type gameView struct {
	scene objects.ObjectInterface
}

func (view *gameView) paint(stdscr *goncurses.Window) {
	for _, obj := range view.scene.GetChildsRecursive() {
		stdscr.Println(obj)
	}

}
