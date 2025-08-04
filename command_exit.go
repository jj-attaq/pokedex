package main

import (
	"fmt"
	"os"

	"github.com/jj-attaq/pokedex/internal"
)

func commandExit(c *internal.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
