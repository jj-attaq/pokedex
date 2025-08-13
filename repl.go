package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jj-attaq/pokedex/internal/pokeapi"
	"github.com/jj-attaq/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
	caughtPokemon    map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

var commands = make(map[string]cliCommand)

func registerCommand(name, desc string, cb func(cfg *config, args ...string) error) {
	commands[name] = cliCommand{
		name:        name,
		description: desc,
		callback:    cb,
	}
}

func init() {
	registerCommand("exit", "Exit the Pokedex", commandExit)
	registerCommand("help", "Displays a help message", commandHelp)
	registerCommand("map", "Lists the next 20 locations", commandMap)
	registerCommand("mapb", "Lists the previous 20 locations", commandMapBack)
	registerCommand("explore", "Lists encounterable Pokemon in specified location", commandExplore)
	registerCommand("catch", "Attempt catching a Pokemon by throwing a pokeball at it", commandCatch)
	registerCommand("inspect", "Provides information on caught Pokemon", commandInspect)
}

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		cmd, exists := commands[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	var result []string
	output := strings.ToLower(text)
	result = strings.Fields(output)

	return result
}
