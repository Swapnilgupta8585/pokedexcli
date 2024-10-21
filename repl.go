package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
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
