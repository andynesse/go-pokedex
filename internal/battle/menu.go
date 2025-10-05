package battle

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func (gs *gameState) eventHandler() {
	ev := gs.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyRight:
			gs.keyRight()
		case tcell.KeyLeft:
			gs.keyLeft()
		case tcell.KeyEnter:
			gs.keyEnter()
		case tcell.KeyEsc:
			gs.running = false
			return
		default:
			r := ev.Rune()
			if r == 'q' || r == 'Q' {
				gs.running = false
				return
			}
		}
	case *tcell.EventResize:
		gs.screen.Sync()
	}
}

func (gs *gameState) keyRight() {
	if gs.mode == "main" {
		gs.selected = min(gs.selected+1, len(gs.menu)-1)
		return
	}
	gs.subSelected = min(gs.subSelected+1, len(gs.options[gs.menu[gs.selected]]))
}

func (gs *gameState) keyLeft() {
	if gs.mode == "main" {
		gs.selected = max(gs.selected-1, 0)
		return
	}
	gs.subSelected = max(gs.subSelected-1, 0)
}

func (gs *gameState) keyEnter() {
	if gs.mode == "main" {
		// enter submenu if items exist
		key := gs.menu[gs.selected]
		if len(gs.options[key]) > 0 {
			gs.mode = "submenu"
			gs.subSelected = 0
		} else {
			// no submenu: perform action (e.g., run)
			fmt.Printf("Action: %s\n", key)
			gs.running = false
		}
	} else {
		// perform submenu action or go back
		topKey := gs.menu[gs.selected]
		items := append([]string{}, gs.options[topKey]...)
		items = append(items, "Back")
		choice := items[gs.subSelected]
		if choice == "Back" {
			gs.mode = "main"
		} else {
			// perform move/item selection
			fmt.Printf("Used %s: %s\n", topKey, choice)
			gs.running = false
		}
	}
}
