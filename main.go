package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		command := cleanInput(scanner.Text())

		if len(command) == 0 {
			continue
		}

		cmd, ok := getCommands()[command[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}