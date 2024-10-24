package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		parts := getCommand(text)

		if len(parts) == 0 {
			continue
		}

		commandName := parts[0]
		args := ""

		if len(parts) > 1 {
			args = parts[1]
		}

		allCommands := commands()
		command, exists := allCommands[commandName]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}

		err := command.callback(&config, args)
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %s\n", err)
	}
}

func getCommand(text string) []string {
	loweredText := strings.ToLower(text)
	fields := strings.Fields(loweredText)
	return fields
}
