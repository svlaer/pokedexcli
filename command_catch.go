package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing argument for \"catch\" command! See \"help\" for more details.")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	respP, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// Using a logistic function to determine the catch-rate based on the Base XP
	baseXP := respP.BaseExperience
	fmt.Printf("*** Base XP for %s: %d ***\n", name, baseXP)
	midpoint := 220
	growthRate := 0.015

	catchRate := 1 - (1 / (1 + math.Exp(-growthRate*(float64(baseXP)-float64(midpoint)))))
	fmt.Printf("*** Catch rate: %f.2 ***\n", catchRate)

	success := (rand.Intn(100) % 100) < int(math.Round(catchRate*100))

	if !success {
		fmt.Printf("%s escaped!\n", name)
		fmt.Println()
	} else {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		fmt.Println()

		pokestats := make(map[string]int)
		for _, val := range respP.Stats {
			pokestats[val.Stat.Name] = val.BaseStat
		}

		poketypes := []string{}
		for _, val := range respP.Types {
			poketypes = append(poketypes, val.Type.Name)
		}

		cfg.pokedex[name] = Pokemon{
			Name:   respP.Name,
			Height: respP.Height,
			Weight: respP.Weight,
			Stats:  pokestats,
			Types:  poketypes,
		}
	}

	return nil
}
