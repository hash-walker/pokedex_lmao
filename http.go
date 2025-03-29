package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func unMarsh() ([]locationArea, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/1/")

	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// parse response to the struct
	var location locationArea
	if err := json.Unmarshal(data, &location); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	output, err := json.MarshalIndent(location, "", "  ")

	if err != nil {
		log.Fatalf("Failed to marshal structured output: %v", err)
	}

	fmt.Printf("%s\n", output)

	return nil, err

}
