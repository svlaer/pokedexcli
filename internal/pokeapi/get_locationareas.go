package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Refactor (see GetLocationArea)
func (c *Client) GetLocationAreas(pageURL *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	data, cacheHit := c.cache.Get(url)
	if !cacheHit {
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

		bodyData, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationAreas{}, fmt.Errorf("Failed to read response body! -> %w", err)
		}

		data = bodyData
		c.cache.Add(url, data)
	}

	var respLA RespLocationAreas
	err := json.Unmarshal(data, &respLA)
	if err != nil {
		return RespLocationAreas{}, fmt.Errorf("Failed to unmarshal Location Areas! -> %w", err)
	}

	return respLA, nil
}
