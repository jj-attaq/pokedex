package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetails(loc string) (Location, error) {
	url := baseURL + "/location-area/" + loc

	if cachedData, found := c.cache.Get(url); found {
		locationResp := Location{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	// cache miss, make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// do and close
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	// jsonify response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// Parse and return struct
	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	// Store data into cache
	c.cache.Add(url, data)

	return locationResp, nil
}
