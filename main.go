package main

import (
	"bufio"
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println("Welcome to Pokedex")
	fmt.Println("Usage: ")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	var commandMap = map[string]cliCommand{

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		command, exists := commandMap[text]
		if !exists {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %s\n", err)
	}

}
