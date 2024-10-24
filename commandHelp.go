package main

import (
	"fmt"
)

func commandHelp(cfg *Config, area string) error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex")
	fmt.Println("Usage: ")
	fmt.Println()
	allCommandsMap := commands()
	for _, value := range allCommandsMap {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}
