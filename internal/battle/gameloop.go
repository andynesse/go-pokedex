package battle

import (
	"fmt"

	"github.com/andynesse/go-pokedex/internal/pokedex"
	"github.com/andynesse/go-pokedex/pkg/config"
	"github.com/gdamore/tcell/v2"
)

type gameState struct {
	Menu        []string
	Options     map[string][]string
	Selected    int
	Mode        string // "top" or "submenu"
	SubSelected int
}

// GameLoop shows a simple four-option selector and lets the user move a
// "*" marker left and right with the arrow keys. Enter will select the
// current option (printed to stdout) and Escape will exit the loop.
func GameLoop(config *config.Config, enemy pokedex.Pokemon) {
	gs := gameState{
		Menu: []string{"attack", "catch", "swap", "bag", "run"},
		Options: map[string][]string{
			"attack": {"Tackle", "Quick Attack", "Thunder Shock", "Tail Whip"},
			"catch":  {"Pokeball", "Greatball", "Ultraball", "Masterball"},
			"swap":   {"Pikachu", "Bulbasaur", "Charmander", "Squirtle"},
			"bag":    {"Potion", "Super Potion", "Hyper Potion", "Max Potion"},
			"run":    {},
		},
		Selected:    0,
		Mode:        "main",
		SubSelected: 0,
	}
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("failed to create screen:", err)
		return
	}
	if err := screen.Init(); err != nil {
		fmt.Println("failed to initialize screen:", err)
		return
	}
	defer screen.Fini()

	draw := func() {
		screen.Clear()
		// Title
		putStr(screen, 0, 0, tcell.StyleDefault.Bold(true), "Battle: "+enemy.Name)

		switch gs.Mode {
		// Render top-level options on one line
		case "main":
			x := 0
			y := 2
			for i, key := range gs.Menu {
				opt := key
				prefix := "  "
				if i == gs.Selected {
					prefix = "* "
				}
				putStr(screen, x, y, tcell.StyleDefault, prefix+opt)
				x += len(prefix) + len(opt) + 4
			}
		case "submenu":
			// Render submenu for the selected top option
			topKey := gs.Menu[gs.Selected]
			items := append([]string{}, gs.Options[topKey]...)
			// add Back option
			items = append(items, "Back")

			y := 2
			putStr(screen, 0, y-1, tcell.StyleDefault.Underline(true), "Choose:")
			x := 0
			for i, it := range items {
				prefix := "  "
				if i == gs.SubSelected {
					prefix = "* "
				}
				putStr(screen, x, y, tcell.StyleDefault, prefix+it)
				x += len(prefix) + len(it) + 4
			}
		}

		screen.Show()
	}

	draw()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRight:
				if gs.Mode == "main" {
					if gs.Selected < len(gs.Menu)-1 {
						gs.Selected++
					}
				} else {
					topKey := gs.Menu[gs.Selected]
					items := append([]string{}, gs.Options[topKey]...)
					items = append(items, "Back")
					if gs.SubSelected < len(items)-1 {
						gs.SubSelected++
					}
				}
				draw()
			case tcell.KeyLeft:
				if gs.Mode == "main" {
					if gs.Selected > 0 {
						gs.Selected--
					}
				} else {
					if gs.SubSelected > 0 {
						gs.SubSelected--
					}
				}
				draw()
			case tcell.KeyEnter:
				if gs.Mode == "main" {
					// enter submenu if items exist
					key := gs.Menu[gs.Selected]
					if len(gs.Options[key]) > 0 {
						gs.Mode = "submenu"
						gs.SubSelected = 0
					} else {
						// no submenu: perform action (e.g., run)
						screen.Fini()
						fmt.Printf("Action: %s\n", key)
						return
					}
				} else {
					// perform submenu action or go back
					topKey := gs.Menu[gs.Selected]
					items := append([]string{}, gs.Options[topKey]...)
					items = append(items, "Back")
					choice := items[gs.SubSelected]
					if choice == "Back" {
						gs.Mode = "main"
					} else {
						// perform move/item selection
						screen.Fini()
						fmt.Printf("Used %s: %s\n", topKey, choice)
						return
					}
				}
				draw()
			case tcell.KeyEsc:
				screen.Fini()
				return
			default:
				r := ev.Rune()
				if r == 'q' || r == 'Q' {
					screen.Fini()
					return
				}
			}
		case *tcell.EventResize:
			screen.Sync()
			draw()
		}
	}
}

// putStr writes s at x,y using the provided style.
func putStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, ch := range str {
		s.SetContent(x, y, ch, nil, style)
		x++
	}
}
