package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		fmt.Println("this command takes no arguments")
	}
	//TODO: see about implementing management commands with this, maybe make it able to clear the dex etc
	fmt.Println("Your Pokemon:")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
