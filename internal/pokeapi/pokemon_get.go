package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if cachedData, found := c.cache.Get(url); found {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(cachedData, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// cache miss, make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// do and close
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// jsonify response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Parse and return struct
	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	// Store data into cache
	c.cache.Add(url, data)

	return pokemonResp, nil
}
