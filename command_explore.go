package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing argument for \"explore\" command! See \"help\" for more details.")
	}
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)

	respLA, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}

	fmt.Println("Found pokemon:")
	for _, result := range respLA.PokemonEncounters {
		fmt.Printf("- %s\n", result.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
