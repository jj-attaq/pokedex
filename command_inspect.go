package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please specify the pokemon you wish to inspect")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("You have not caught %s yet.", name)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
