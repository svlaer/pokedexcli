package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, prs := commands()[commandName]

		if prs {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Unrecognized command: ", commandName)
			fmt.Println("Try the \"help\" command for more info.")
			fmt.Println()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Reading standard input: ", err)
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
	callback    func() error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
