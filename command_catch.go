package main

import (
	"fmt"
	"math/rand"
)

func chance(n int) bool {
	chance := rand.Intn(700)
	return chance >= n
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	// run pokeapi lookup of pokemon in question
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience

	// catching logic

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	wasCaught := chance(baseExp)
	if wasCaught {
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
