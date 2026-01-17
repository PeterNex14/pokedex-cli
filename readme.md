# Pokedex CLI

## Description
This project is a CLI (Command Line Interface) Pokedex application written in Go. It interacts with the [PokeAPI](https://pokeapi.co/) to simulate the experience of exploring the Pokemon world, finding Pokemon in different areas, and catching them.

## What I've been working on
I have been building a REPL (Read-Eval-Print Loop) in Go that processes user commands to interact with internet APIs. The application handles:
- Parsing user input.
- Making HTTP requests to the PokeAPI.
- Parsing JSON responses into Go structs.
- Implementing a custom caching system to reduce network calls.
- Game logic for catching Pokemon based on their base experience stats.

## What I've been learning
This project is part of the guided project on [Boot.dev](https://boot.dev). Through this process, I have learned/practiced:
- **Go Syntax & Structure**: Using structs, slices, maps, and loops.
- **HTTP Client**: Sending requests and handling responses using `net/http`.
- **JSON Handling**: Unmarshalling JSON data into Go structures.
- **CLI Architecture**: structure a basic command-line tool with a command registry.
- **Caching**: Implementing a basic caching mechanism to optimize API usage.

## Usage
Run the program:
```bash
go run .
```

Available commands:
- `help`: Displays a help message
- `exit`: Exit the Pokedex
- `map`: Displays the names of 20 location areas in the Pokemon world
- `mapb`: Displays the previous 20 locations areas
- `explore <area_name>`: List all of the pokemon based on the specified location
- `catch <pokemon_name>`: Attempt to catch a Pokemon
- `inspect <pokemon_name>`: View details of a caught Pokemon
- `pokedex`: List all caught Pokemon

## Future Improvements & Ideas
- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
