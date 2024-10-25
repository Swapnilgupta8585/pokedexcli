package main

import (
	"fmt"
)

func commandInspect(cfg *Config, pokemonName string) error {
	pokemon, ok := cfg.user[pokemonName]
	if !ok {
		fmt.Println()
		fmt.Println("you have not caught that pokemon")
		fmt.Println()
		return nil
	}
	fmt.Println()
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stats := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types: ")
	for _, types := range pokemon.Types {
		fmt.Printf("  -%s\n", types.Type.Name)
	}
	fmt.Println()
	return nil
}
