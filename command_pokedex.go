package main

import (
	"fmt"
	"maps"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for k := range maps.Keys(cfg.pokedex) {
		fmt.Printf("\t- %s\n", k)
	}
	fmt.Println()

	return nil
}
