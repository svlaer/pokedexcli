package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMap(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("Failed to perform request for \"GetMap\"! -> %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreas{}, fmt.Errorf("Failed to get Location Areas! Status code: %s -- (URL: %s)", res.Status, url)
	}

	var la LocationAreas
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&la); err != nil {
		return LocationAreas{}, fmt.Errorf("Failed to decode Location Areas! -> %w", err)
	}

	return la, nil
}
