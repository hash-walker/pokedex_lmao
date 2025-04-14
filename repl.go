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
	caughtPokemon *map[string]pokeapi.RespPokemon
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
		arguments := []string{}
		if len(words) > 1{
			arguments = words[1:]
		}

		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg, arguments...)

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
		"explore":{
			name: "explore <location_name>",
			description: "Get the pokemons at that location area",
			callback: exploreArea,
		},

		"catch":{
			name: "catch <pokemon_name>",
			description: "Catch the pokemon",
			callback: catchPokemon,
		},

		"inspect":{
			name: "inspect <pokemon_name>",
			description: "Information about caught pokemon",
			callback: inspectPokemon,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
