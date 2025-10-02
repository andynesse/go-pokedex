package commands

import (
	"fmt"

	"github.com/andynesse/go-pokedex/pkg/config"
)

func commandPokedex(config *config.Config) error {
	if len(config.Pokedex.Pokemon) == 0 {
		fmt.Println("no caught pokemon")
		return nil
	}
	output := "Your Pokedex:\n"
	for name := range config.Pokedex.Pokemon {
		output += fmt.Sprintf("  - %s\n", name)
	}
	fmt.Print(output)
	return nil
}
