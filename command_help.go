package main

import (
	"fmt"
	// "github.com/jj-attaq/pokedex/internal"
)

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	fmt.Printf("\n")

	return nil
}
