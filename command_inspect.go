package main

import (
	"errors"
	"fmt"
)

func inspectPokemon(cfg *config, arguments ...string) error{

	if len(arguments) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := arguments[0]


	caught, ok := (*cfg.caughtPokemon)[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf(
		"Name: %s\nHeight: %d\nWeight: %d\n",	caught.Name, caught.Height, caught.Weight,
	)

	fmt.Println("Stats:")
	for _,stats := range caught.Stats{
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Println("Types:")
	for _,types := range caught.Types{
		fmt.Printf("	- %s\n", types.Type.Name)
	}
	

	return nil
}