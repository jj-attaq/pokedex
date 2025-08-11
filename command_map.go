package main

import (
	"fmt"
)

// does the same as ListLocations function in the solution here https://www.boot.dev/solution/813eafe1-2e1d-42a0-b358-53e0f4d4fdc8
// ListLocations has better separation of concerns, while mine is more grugbrained
// func determineUrlPosition(c *config, res *http.Response) (*config, error) {
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("reading response body failed: %w", err)
// 	}
//
// 	if res.StatusCode > 299 {
// 		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
// 	}
//
// 	areas := pokeapi.RespShallowLocations{}
// 	err = json.Unmarshal(body, &areas)
// 	if err != nil {
// 		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
// 	}
//
// 	for _, area := range areas.Results {
// 		fmt.Println(area.Name)
// 	}
// 	c.nextLocationsURL = areas.Next
// 	c.prevLocationsURL = areas.Previous
//
// 	return c, nil
// }

func commandMap(c *config) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
	// var urlStr string
	// if c.nextLocationsURL == nil {
	// 	urlStr = "https://pokeapi.co/api/v2/location-area/"
	// } else {
	// 	urlStr = *c.nextLocationsURL
	// }
	//
	// res, err := c.fetchWithCache(urlStr)
	// if err != nil {
	// 	return fmt.Errorf("HTTP request failed: %w", err)
	// }
	//
	// updatedConf, err := determineUrlPosition(c, res)
	// if err != nil {
	// 	return fmt.Errorf("ERROR: %w", err)
	// }
	//
	// c = updatedConf
	//
	// return nil
}

func commandMapBack(c *config) error {
	if c.prevLocationsURL == nil {
		fmt.Errorf("You're on the first page")
	}

	locationsResp, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
	// var urlStr string
	// if c.prevLocationsURL == nil {
	// 	return fmt.Errorf("Cannot go back further.")
	// } else {
	// 	urlStr = *c.prevLocationsURL
	// }
	//
	// res, err := c.fetchWithCache(urlStr)
	// if err != nil {
	// 	return fmt.Errorf("HTTP request failed: %w", err)
	// }
	// defer res.Body.Close()
	//
	// updatedConf, err := determineUrlPosition(c, res)
	// if err != nil {
	// 	return fmt.Errorf("ERROR: %w", err)
	// }
	//
	// c = updatedConf
	//
	// return nil
}
