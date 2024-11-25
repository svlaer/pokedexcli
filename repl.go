package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/svlaer/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, prs := commands()[commandName]

		if prs {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
			}
		} else {
			fmt.Fprintln(os.Stderr, "Unrecognized command: ", commandName)
			fmt.Println("Try the \"help\" command for more info.")
			fmt.Println()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Reading standard input: ", err)
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 names of location areas in the Pokemon world. Each subsequent invocation displays the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "(Map back) Like the \"map\" command. Displays the previous 20 locations.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location-area-name>",
			description: "Displays a list of all Pokemon in a given area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch the given pokemon and add its information to your pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "See the details of a caught pokemon.",
			callback:    commandInspect,
		},
	}
}
