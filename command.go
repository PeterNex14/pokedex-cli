package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name 		string
	description	string
	callback	func(*Config) error
} 

type Config struct {
	Next 		*string
	Previous 	*string
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	locations, err := getLocationArea(url)

	if err != nil {
		return err
	}

	for _, value := range locations.Results {
		fmt.Println(value.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil

}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} 

	url := *cfg.Previous

	locations, err := getLocationArea(url)

	if err != nil {
		return err
	}

	for _, value := range locations.Results {
		fmt.Println(value.Name)
	}

	cfg.Previous = locations.Previous
	cfg.Next = locations.Next

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name: "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations areas in the Pokemon world.",	
			callback: commandMapb,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}



