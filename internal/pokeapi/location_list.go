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

	//get
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

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

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
