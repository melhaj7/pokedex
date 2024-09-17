package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu")
	fmt.Println("Here is a list of available commands")

	commandsMap := getCommands()

	for _, cmd := range commandsMap {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
