package commands

import (
	"fmt"
	"os"

	"github.com/andynesse/go-pokedex/pkg/config"
)

func commandExit(config *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp(config *config.Config) error {
	output := fmt.Sprintln("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range commands {
		output += fmt.Sprintf("\n%s: %s", cmd.Name, cmd.Description)
	}
	fmt.Println(output)
	return nil
}
