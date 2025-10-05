package battle

import (
	"github.com/gdamore/tcell/v2"
)

func (gs *gameState) draw() {
	gs.screen.Clear()
	// Title
	putStr(gs.screen, 0, 0, tcell.StyleDefault.Bold(true), "Battle: "+gs.enemy.Name)

	switch gs.mode {
	// Render top-level options on one line
	case "main":
		x := 0
		y := 2
		for i, key := range gs.menu {
			opt := key
			prefix := "  "
			if i == gs.selected {
				prefix = "* "
			}
			putStr(gs.screen, x, y, tcell.StyleDefault, prefix+opt)
			x += len(prefix) + len(opt) + 4
		}
	case "submenu":
		// Render submenu for the selected top option
		topKey := gs.menu[gs.selected]
		items := append([]string{}, gs.options[topKey]...)
		// add Back option
		items = append(items, "Back")

		y := 2
		putStr(gs.screen, 0, y-1, tcell.StyleDefault.Underline(true), "Choose:")
		x := 0
		for i, it := range items {
			prefix := "  "
			if i == gs.subSelected {
				prefix = "* "
			}
			putStr(gs.screen, x, y, tcell.StyleDefault, prefix+it)
			x += len(prefix) + len(it) + 4
		}
	}

	gs.screen.Show()
}
