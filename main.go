package main

import (
	"github.com/hash-walker/pokedex_lmao/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)

	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)

}
