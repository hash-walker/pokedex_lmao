package main

import "fmt"

func pokedex(cfg *config, argument ...string) error {

	if cfg.caughtPokemon == nil {
		fmt.Println("No pokemon are caught yet")
		return nil
	}

	fmt.Println("Your Pokedex")
	for _, pokemon := range (*cfg.caughtPokemon) {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}