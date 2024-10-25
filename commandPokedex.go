package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, s string) error {
	if len(cfg.user) == 0 {
		fmt.Println()
		fmt.Println("You haven't caught any pokemon!")
		fmt.Println()
		return nil
	}
	fmt.Println()
	for pokemonCaught, _ := range cfg.user {
		fmt.Printf("Your Pokedex:\n")
		fmt.Printf(" -%s\n", pokemonCaught)
	}
	fmt.Println()
	return nil
}
