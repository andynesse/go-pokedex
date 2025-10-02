package repl

import (
	"fmt"
	"strings"

	"github.com/andynesse/go-pokedex/internal/commands"
	"github.com/andynesse/go-pokedex/pkg/config"
	"github.com/chzyer/readline"
)

func Run(config *config.Config) {
	rl, err := readline.New("Pokedex > ")
	if err != nil {
		fmt.Println("Failed to initialize readline:", err)
		return
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		input := cleanInput(line)
		cmdInput := input[0]
		config.Args = input[1:]
		command, exists := commands.GetCommand(cmdInput)
		if !exists {
			fmt.Printf("Unknown command: %s\n", cmdInput)
			continue
		}
		if err := command.Callback(config); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(text, "\t", " "), "\n", " "))
	if len(trimmed) == 0 {
		return []string{}
	}
	splitted := strings.Split(strings.ToLower(trimmed), " ")
	output := []string{}
	for _, word := range splitted {
		if len(word) != 0 {
			output = append(output, word)
		}
	}
	return output
}
