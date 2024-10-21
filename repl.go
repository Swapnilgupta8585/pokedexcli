package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var config = Config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: nil,
	}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		allCommands := commands()
		command, exists := allCommands[text]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}

		err := command.callback(&config)
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %s\n", err)
	}
}
