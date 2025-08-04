package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jj-attaq/pokedex/internal"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func determineUrlPosition(c *internal.Config, res *http.Response) (*internal.Config, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body failed: %w", err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	areas := Locations{}
	err = json.Unmarshal(body, &areas)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	c.Next = areas.Next
	c.Previous = areas.Previous

	return c, nil
}

func commandMap(c *internal.Config) error {
	url := c.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer res.Body.Close()

	updatedConf, err := determineUrlPosition(c, res)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	c = updatedConf

	return nil
}

func commandMapBack(c *internal.Config) error {
	url := c.Previous
	if url == "" {
		return fmt.Errorf("Cannot go back further.")
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer res.Body.Close()

	updatedConf, err := determineUrlPosition(c, res)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	c = updatedConf

	return nil
}
