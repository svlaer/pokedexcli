package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing argument for \"inspect\" command! See \"help\" for more details.")
	}
	name := args[0]

	poke, prs := cfg.pokedex[name]
	if !prs {
		fmt.Println("You haven't caught that Pokemon.")
		fmt.Println()
		return nil
	}

	fmt.Println(poke)
	fmt.Println()

	return nil
}
