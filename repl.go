package main

import (
	"bufio"
	"fmt"
	"github.com/hash-walker/pokedex_lmao/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("\nGoodbye!")
			break
		}

		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg)

			if err != nil {
				fmt.Println(err)
			}

		}

	}
}

func getCommands() map[string]cliCommand {
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
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
