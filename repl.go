package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jj-attaq/pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *internal.Config) error
}

var commands = make(map[string]cliCommand)

func registerCommand(name, desc string, cb func(c *internal.Config) error) {
	commands[name] = cliCommand{
		name:        name,
		description: desc,
		callback:    cb,
	}
}

func init() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
	registerCommand("map", "lists the next 20 locations", commandMap)
	registerCommand("mapb", "lists the previous 20 locations", commandMapBack)
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := internal.Config{}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, exists := commands[commandName]
		if exists {
			err := cmd.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

		// input := CleanInput(scanner.Text())
		// if len(input) > 1 {
		// 	fmt.Printf("Unknown command: %s. Commands must be a single word.\n", input)
		// 	continue
		// }
		// inputCmd := input[0]
		//
		// cmd := commands[inputCmd]
		// if inputCmd != cmd.name {
		// 	fmt.Printf("Unknown command: %s\n", input)
		// } else {
		// 	cmd.callback()
		// }
	}
}

func CleanInput(text string) []string {
	var result []string
	output := strings.ToLower(text)
	result = strings.Fields(output)

	return result
}
