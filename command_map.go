package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, argument ...string) error {

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, argumemnt ...string) error {

	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous // Correctly update previousLocationsURL

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil

}
