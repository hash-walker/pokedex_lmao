package main

import (
	"fmt"
	"errors"
)

func exploreArea(cfg *config, arguments ...string) error{

	if len(arguments) != 1 {
		return errors.New("you must provide a location name")
	}

	location := arguments[0]

	fmt.Printf("Exploring %s...\n", location)

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(location)

	if err != nil {
		return err
	}

	pokemonEncounters := pokemonResp.PokemonEncounters

	
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range pokemonEncounters{
		fmt.Printf("	- %s\n", pokemonEncounter.Pokemon.Name)
	}


	return nil
}