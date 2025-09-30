package commands

import "github.com/andynesse/go-pokedex/internal/config"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
}

var commands map[string]cliCommand

func init() {
	commands = make(map[string]cliCommand)
	commands["exit"] = cliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	}
	commands["help"] = cliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		Name:        "map",
		Description: "Displays 20 next cities",
		Callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		Name:        "mapb",
		Description: "Displays 20 previous cities",
		Callback:    commandMapback,
	}
	commands["explore"] = cliCommand{
		Name:        "explore",
		Description: "Displays a list of pokemon located at the given area",
		Callback:    commandExplore,
	}
}

func GetCommand(key string) (cliCommand, bool) {
	command, exists := commands[key]
	if !exists {
		return cliCommand{}, false
	}
	return command, true
}
