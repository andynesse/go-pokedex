package battle

import (
	"fmt"

	"github.com/andynesse/go-pokedex/internal/pokedex"
	"github.com/andynesse/go-pokedex/pkg/config"
	"github.com/gdamore/tcell/v2"
)

type gameState struct {
	menu        []string
	options     map[string][]string
	selected    int
	mode        string
	subSelected int
	screen      tcell.Screen
	enemy       pokedex.Pokemon
	running     bool
}

func Start(config *config.Config, enemy pokedex.Pokemon) {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println("failed to create screen:", err)
		return
	}
	if err := screen.Init(); err != nil {
		fmt.Println("failed to initialize screen:", err)
		return
	}
	gs := gameState{
		menu: []string{"attack", "catch", "swap", "bag", "run"},
		options: map[string][]string{
			"attack": {"Tackle", "Quick Attack", "Thunder Shock", "Tail Whip"},
			"catch":  {"Pokeball", "Greatball", "Ultraball", "Masterball"},
			"swap":   {"Pikachu", "Bulbasaur", "Charmander", "Squirtle"},
			"bag":    {"Potion", "Super Potion", "Hyper Potion", "Max Potion"},
			"run":    {},
		},
		selected:    0,
		mode:        "main",
		subSelected: 0,
		screen:      screen,
		enemy:       enemy,
		running:     true,
	}
	gameLoop(gs)
}
func gameLoop(gs gameState) {

	defer gs.screen.Fini()

	for gs.running {
		gs.draw()
		gs.eventHandler()
	}
}

func putStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, ch := range str {
		s.SetContent(x, y, ch, nil, style)
		x++
	}
}
