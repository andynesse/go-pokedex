package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init() {
	commands = make(map[string]cliCommand)
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			continue
		}
		input := cleanInput(scanner.Text())
		command, ok := commands[input[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", input[0])
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	output := fmt.Sprintln("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range commands {
		output += fmt.Sprintf("\n%s: %s", cmd.name, cmd.description)
	}
	fmt.Println(output)
	return nil
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
