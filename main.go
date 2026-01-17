package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		pokedex: make(map[string]Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		command := cleanInput(scanner.Text())

		if len(command) == 0 {
			continue
		}

		cmdName := command[0]
		args := []string{}

		if len(command) > 1 {
			args = command[1:]
		}

		cmd, ok := getCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
		} 

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}