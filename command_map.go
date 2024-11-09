package main

import (
	"fmt"

	"github.com/svlaer/pokedexcli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	url := config.Next

	la, err := pokeapi.GetMap(url)
	if err != nil {
		return err
	}

	config.Next = la.Next
	if la.Previous == nil {
		config.Previous = ""
	} else {
		config.Previous = *la.Previous
	}

	for _, result := range la.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()

	return nil
}
