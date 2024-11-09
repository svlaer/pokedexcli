package main

import (
	"os"

	"github.com/svlaer/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.Config) error {
	os.Exit(0)
	return nil
}
