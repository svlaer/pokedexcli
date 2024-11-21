package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Refactor to general PerformGetWithCaching (or GetResource or something) function, also for GetLocationAreas and GetPokemon
func (c *Client) GetLocationArea(locationAreaName string) (RespLocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s/", baseURL, locationAreaName)

	data, cacheHit := c.cache.Get(url)
	if !cacheHit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationArea{}, fmt.Errorf("Failed to build request for \"GetLocationArea\"! -> %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationArea{}, fmt.Errorf("Failed to perform request for \"GetLocationArea\"! -> %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return RespLocationArea{}, fmt.Errorf("Failed to get Location Area! Status code: %s -- (URL: %s)", res.Status, url)
		}

		bodyData, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationArea{}, fmt.Errorf("Failed to read response body! -> %w", err)
		}

		data = bodyData
		c.cache.Add(url, data)
	}

	var respLA RespLocationArea
	err := json.Unmarshal(data, &respLA)
	if err != nil {
		return RespLocationArea{}, fmt.Errorf("Failed to unmarshal Location Area! -> %w", err)
	}

	return respLA, nil
}
