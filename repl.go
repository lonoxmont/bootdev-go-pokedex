package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lonoxmont/bootdev-go-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	consolescanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		consolescanner.Scan()
		input := consolescanner.Text()
		if len(input) == 0 {
			continue
		}
		words := cleanInput(input)
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try to catch the supplied Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
