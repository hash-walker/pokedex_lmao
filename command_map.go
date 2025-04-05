package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {

	fmt.Printf("hello, %v\n", cfg.nextLocationsURL)
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	fmt.Print(*locationResp.Next)

	cfg.nextLocationsURL = locationResp.Next
	cfg.nextLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {

	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.nextLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil

}
