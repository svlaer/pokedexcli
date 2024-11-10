package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, fmt.Errorf("Failed to build request for \"GetLocationAreas\"! -> %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, fmt.Errorf("Failed to perform request for \"GetLocationAreas\"! -> %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return RespLocationAreas{}, fmt.Errorf("Failed to get Location Areas! Status code: %s -- (URL: %s)", res.Status, url)
	}

	var respLA RespLocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&respLA); err != nil {
		return RespLocationAreas{}, fmt.Errorf("Failed to decode Location Areas! -> %w", err)
	}

	return respLA, nil
}
