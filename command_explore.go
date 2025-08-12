package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}

	location := args[0]
	detailsResp, err := cfg.pokeapiClient.GetLocationDetails(location)
	if err != nil {
		return nil
	}

	fmt.Printf("Exploring %s...\n", detailsResp.Location.Name)
	fmt.Printf("Found Pokemon:\n")

	for _, pokemon := range detailsResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
