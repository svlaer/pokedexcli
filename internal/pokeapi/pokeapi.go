package pokeapi

type Config struct {
	Next     string
	Previous string
}

func InitialConfig() Config {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	return Config{
		Next:     baseUrl,
		Previous: "",
	}
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
