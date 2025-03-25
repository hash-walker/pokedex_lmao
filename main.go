package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		fmt.Printf("Your command was: %s\n", words[0])
		fmt.Print("Pokedex > ")
	}
}
