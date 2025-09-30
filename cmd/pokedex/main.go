package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/andynesse/go-pokedex/internal/commands"
	"github.com/andynesse/go-pokedex/internal/config"
	"github.com/andynesse/go-pokedex/internal/pokecache"
	"github.com/andynesse/go-pokedex/internal/pokedex"
	"github.com/andynesse/go-pokedex/internal/repl"
)

func main() {
	config := config.Config{
		Next:    "https://pokeapi.co/api/v2/location-area/",
		Cashe:   pokecache.NewCache(5 * time.Second),
		Pokedex: *pokedex.NewPokedex(),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			continue
		}
		input := repl.CleanInput(scanner.Text())
		cmdInput := input[0]
		config.Args = input[1:]
		command, exists := commands.GetCommand(cmdInput)
		if !exists {
			fmt.Printf("Unknown command: %s\n", cmdInput)
			continue
		}
		if err := command.Callback(&config); err != nil {
			fmt.Println(err)
		}
	}
}
