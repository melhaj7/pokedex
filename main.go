package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func main() {
	commands := make(map[string]cliCommand)
	url := "https://pokeapi.co/api/v2/location-area/"

	locationResponse, err := GetLocation(url)
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}

	fmt.Println("First 20 locations: ")
	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    func() { commandHelp(commands) },
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())

			cmd, exist := commands[input]
			if exist {
				cmd.callback()
			} else {
				fmt.Println("Unknown command. Type 'help' for a list of commands")
			}
		}
	}

}

func commandHelp(commands map[string]cliCommand) {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
}

func commandExit() {
	fmt.Println("Exiting the Pokedex. Goodbye!")
	os.Exit(0)
}
