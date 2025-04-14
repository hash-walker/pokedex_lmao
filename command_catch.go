package main

import (
	"math/rand"
	"errors"
	"fmt"

	"github.com/hash-walker/pokedex_lmao/internal/pokeapi"
)


func catchPokemon(cfg *config, arguments ...string) error {

	if cfg.caughtPokemon == nil {
		cfg.caughtPokemon = &map[string]pokeapi.RespPokemon{}
	}

	if len(arguments) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := arguments[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemonInfo, err := cfg.pokeapiClient.Pokemon(name)

	if err != nil {
		return err
	}
	
	catch := rand.Intn(2)


	if catch == 1 {
		fmt.Printf("%s was caught\n", pokemonInfo.Name)

		if (*cfg.caughtPokemon) == nil {
			(*cfg.caughtPokemon) = make(map[string]pokeapi.RespPokemon)
		}

		caught := *cfg.caughtPokemon
		caught[name] = pokemonInfo

	}else{
		fmt.Printf("%s escaped\n", pokemonInfo.Name)
	}

	return nil
}