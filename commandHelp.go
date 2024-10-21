package main

import (
	"fmt"
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
