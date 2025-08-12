package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetails(location string) (RespAreaDetails, error) {
	url := baseURL + "/location-area/" + location

	if cachedData, found := c.cache.Get(url); found {
		detailsResp := RespAreaDetails{}
		err := json.Unmarshal(cachedData, &detailsResp)
		if err != nil {
			return RespAreaDetails{}, err
		}
		return detailsResp, nil
	}

	// cache miss, make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaDetails{}, err
	}

	// do and close
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaDetails{}, err
	}
	defer resp.Body.Close()

	// jsonify response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespAreaDetails{}, err
	}

	// Parse and return struct
	detailsResp := RespAreaDetails{}
	err = json.Unmarshal(data, &detailsResp)
	if err != nil {
		return RespAreaDetails{}, err
	}

	// Store data into cache
	c.cache.Add(url, data)

	return detailsResp, nil
}
