package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, found := c.cache.Get(url); found {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	// cache miss, make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	//do and close
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	//jsonify response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Parse and return struct
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Store data into cache
	c.cache.Add(url, data)

	return locationsResp, nil
}

// func (c *config) fetchWithCache(url string) ([]byte, error) {
// 	// Check cache
// 	if cachedData, found := c.cache.Get(url); found {
// 		return cachedData, nil
// 	}
//
// 	// Cache miss, make request
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()
//
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Store in cache
// 	c.cache.Add(url, body)
//
// 	return body, nil
// }
