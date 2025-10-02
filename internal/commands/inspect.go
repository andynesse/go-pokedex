package commands

import (
	"fmt"

	"github.com/andynesse/go-pokedex/pkg/config"
)

func commandInspect(config *config.Config) error {
	if len(config.Args) == 0 {
		return fmt.Errorf("no pokemon to inspect")
	}
	pokemon, exists := config.Pokedex.Pokemon[config.Args[0]]
	if !exists {
		return fmt.Errorf("pokemon does not exist in pokedex")
	}

	output := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for k, v := range pokemon.Stats {
		output += fmt.Sprintf("  -%s: %d\n", k, v)
	}
	output += "Types:\n"
	for _, t := range pokemon.Types {
		output += fmt.Sprintf("  - %s\n", t)
	}
	fmt.Print(output)
	return nil
}
