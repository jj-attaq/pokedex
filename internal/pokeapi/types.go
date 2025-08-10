package pokeapi

//	type Config struct {
//		Next     string `json:"next"`
//		Current  string `json:"current"`
//		Previous string `json:"previous"`
//	}
//
// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
