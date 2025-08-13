package main

import (
	"time"

	"github.com/jj-attaq/pokedex/internal/pokeapi"
	"github.com/jj-attaq/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	pokeClient := pokeapi.NewClient(5*time.Minute, cache)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	StartRepl(cfg)
}
