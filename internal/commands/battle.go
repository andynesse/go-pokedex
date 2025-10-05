package commands

import (
	"fmt"

	"github.com/andynesse/go-pokedex/internal/battle"
	"github.com/andynesse/go-pokedex/internal/pokedex"
	"github.com/andynesse/go-pokedex/pkg/config"
)

func commandBattle(config *config.Config) error {
	pokemon := pokedex.Pokemon{
		Name:   "machop",
		Weight: 10,
		Height: 10,
		Stats: map[string]int{
			"speed":           50,
			"hp":              50,
			"attack":          50,
			"defence":         50,
			"special_attack":  50,
			"special_defence": 50,
		},
		Types: []string{
			"fighting",
		},
	}
	battle.GameLoop(config, pokemon)
	return fmt.Errorf("battle ended")
}
