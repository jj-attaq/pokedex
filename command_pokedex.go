package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.caughtPokemon
	if len(pokedex) < 1 {
		return fmt.Errorf("Your Pokedex is empty.")
	}

	fmt.Printf("Your Pokedex:\n")
	for _, val := range pokedex {
		fmt.Printf(" - %s\n", val.Name)
	}
	return nil
}
