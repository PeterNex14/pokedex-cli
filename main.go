package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}