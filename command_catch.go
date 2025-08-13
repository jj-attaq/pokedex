package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jj-attaq/pokedex/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.Pokemon)

func chance(n int) bool {
	// // Convert to non-negative value (absolute value)
	// base := n
	// if base < 0 {
	// 	base = -base
	// }
	//
	// // Handle potential integer overflow edge case
	// if base == (1<<63 - 1) { // Max int64 value
	// 	// For maximum integer value, use 50% probability
	// 	return rand.Intn(2) == 0
	// }
	//
	// // Generate random number in [0, base] range
	// // Return true if we get 0 (probability = 1/(base+1))
	// return rand.Intn(base+1) == 0
	chance := rand.Intn(700)
	return chance >= n
}

func commandCatch(cfg *config, args ...string) error {
	// run pokeapi lookup of pokemon in question
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	pokemon := pokemonResp.Name
	baseExp := pokemonResp.BaseExperience

	// seed
	rand.NewSource(time.Now().UnixNano())

	// catching logic

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	wasCaught := chance(baseExp)
	if wasCaught {
		fmt.Printf("%s was caught!\n", pokemon)
		pokedex[pokemon] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	for key, _ := range pokedex {
		fmt.Printf(" - %s\n", pokedex[key].Name)
	}

	return nil
}
