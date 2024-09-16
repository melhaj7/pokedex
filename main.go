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

var config = &pokeapi.Config{
	Next:     nil,
	Previous: nil,
	Current:  nil,
}

func main() {
	// commands := make(map[string]cliCommand)
	baseURL := "https://pokeapi.co/api/v2/location-area/"

	// commands["help"] = cliCommand{
	// 	name:        "help",
	// 	description: "Displays a help message",
	// 	callback:    func() { commandHelp(commands) },
	// }
	// commands["exit"] = cliCommand{
	// 	name:        "exit",
	// 	description: "Exit the pokedex",
	// 	callback:    commandExit,
	// }
	// commands["map"] = cliCommand{
	// 	name:        "map",
	// 	description: "fetch 20 locations from pokeapi",
	// 	callback:    getAndPrintLocations(config.Next, baseURL),
	// }
	// commands["mapb"] = cliCommand{
	// 	name:        "map",
	// 	description: "fetch 20 locations from pokeapi",
	// 	callback:    getAndPrintLocations(config.Previous, baseURL),
	// }

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())

			// cmd, exist := commands[input]
			// if exist {
			// 	cmd.callback()
			// } else {
			// 	fmt.Println("Unknown command. Type 'help' for a list of commands")
			// }
			switch input {
			case "help":
				commandHelp()
			case "exit":
				commandExit()
			case "map":
				getAndPrintLocations(config.Next, baseURL, "next")
			case "mapb":
				getAndPrintLocations(config.Previous, baseURL, "previous")
			default:
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
			}
		}
	}

}

// commands map[string]cliCommand
func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	// for _, cmd := range commands {
	// 	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	// }
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Display the next 20 location areas")
	fmt.Println("mapb: Display the previous 20 location areas")
}

func commandExit() {
	fmt.Println("Exiting the Pokedex. Goodbye!")
	os.Exit(0)
}

func getAndPrintLocations(url *string, baseURL string, direction string) {
	if url == nil {
		if direction == "next" {
			url = &baseURL
		} else {
			fmt.Println("No previous locations available.")
			return
		}
	}
	locationResponse, err := pokeapi.GetLocation(baseURL)
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
		return
	}
	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous
	config.Current = url

	fmt.Println("Locations: ")
	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}

	if locationResponse.Next == nil && direction == "next" {
		fmt.Println("No more locations to display.")
	}
	if locationResponse.Previous == nil && direction == "previous" {
		fmt.Println("You are at the start. No previous locations available.")
	}
}
