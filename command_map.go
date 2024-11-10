package main

import (
	"fmt"
	"os"
)

func commandMap(cfg *config) error {
	respLA, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = respLA.Next
	cfg.prevLocationsURL = respLA.Previous

	for _, result := range respLA.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Fprintln(os.Stderr, "First page! Can't map back.")
		fmt.Println()
		return nil
	}

	respLA, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = respLA.Next
	cfg.prevLocationsURL = respLA.Previous

	for _, result := range respLA.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()

	return nil
}
