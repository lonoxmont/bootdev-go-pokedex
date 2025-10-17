package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	consolescanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for consolescanner.Scan() {
		input := consolescanner.Text()
		if len(input) == 0 {
			fmt.Print("Pokedex > ")
			continue

		}
		words := cleanInput(input)
		commandName := words[0]
		fmt.Printf("Your command was: %s\n", commandName)
		fmt.Print("Pokedex > ")

	}

}

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	return words
}
