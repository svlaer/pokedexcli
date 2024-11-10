package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	cmds := commands()

	// Sort the keys so that the help menu has the same order every time
	keySlice := make([]string, 0)
	for k, _ := range cmds {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)

	for _, e := range keySlice {
		cmd := cmds[e]
		fmt.Printf("%s: \n\t%s\n\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}
