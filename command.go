package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
	"github.com/PeterNex14/pokedex-cli/internal/pokecache"
)

type cliCommand struct {
	name 		string
	description	string
	callback	func(*Config, ...string) error
} 

type Config struct {
	Next 		*string
	Previous 	*string
	pokedex 	map[string]Pokemon
}

var cache = pokecache.NewCache(time.Second * 5)

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	locations, err := getLocationArea(url, cache)

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

func commandMapb(cfg *Config, args ...string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	} 

	url := *cfg.Previous

	locations, err := getLocationArea(url, cache)

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

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Argument must be specified")
	}

	commandArgs := args[0]
	baseUrl := "https://pokeapi.co/api/v2/location-area/"

	pokemon, err := getPokemonLocation(commandArgs, baseUrl, cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _,value := range pokemon.PokemonEncounter {
		fmt.Printf("- %s\n", value.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Argument must be specified")
	}

	commandArgs := args[0]
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"

	pokemon, err := getPokemon(commandArgs, baseUrl)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", commandArgs)
	
	num := rand.Intn(100)
	chance := 100 - (pokemon.BaseExperience / 10)

	if num < chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch" : {
			name: "catch",
			description: "Catch Pokemon based on the given name of pokemon",
			callback: commandCatch,
		},
		"explore": {
			name: "explore",
			description: "List all of the pokemon based on the specified location",
			callback: commandExplore,
		},
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



