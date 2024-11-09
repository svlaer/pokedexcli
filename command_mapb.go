package main

import (
	"fmt"
	"os"

	"github.com/svlaer/pokedexcli/internal/pokeapi"
)

func commandMapb(config *pokeapi.Config) error {
	url := config.Previous

	if url == "" {
		fmt.Fprintln(os.Stderr, "First page! Can't map back.")
		fmt.Println()
		return nil
	}

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
