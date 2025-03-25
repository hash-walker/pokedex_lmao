package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		text := scanner.Text()

		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		firstCommand := words[0]

		fmt.Printf("Your command was: %s\n", firstCommand)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
