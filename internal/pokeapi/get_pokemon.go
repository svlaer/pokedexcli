package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Refactor (See GetLocationArea)
func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s/", baseURL, pokemonName)

	data, cacheHit := c.cache.Get(url)
	if !cacheHit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("Failed to build request for \"GetPokemon\"! -> %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("Failed to perform request for \"GetPokemon!\" -> %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return RespPokemon{}, fmt.Errorf("Failed to get Pokemon! Status code: %s -- (URL: %s)", res.Status, url)
		}

		bodyData, err := io.ReadAll(res.Body)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("Failed to read response body! -> %w", err)
		}

		data = bodyData
		c.cache.Add(url, data)
	}

	var respP RespPokemon
	err := json.Unmarshal(data, &respP)
	if err != nil {
		return RespPokemon{}, fmt.Errorf("Failed to unmarshal Pokemon! -> %w", err)
	}

	return respP, nil
}
