package main

import (
	"fmt"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	fmt.Println()
	locations, err := pokeapi.PokeapiCall(cfg.next)
	if err != nil {
		return fmt.Errorf("can not get the locations: %w", err)
	}
	cfg.next = *locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("Pokedex > %s\n", location.Name)
	}
	fmt.Println()
	return nil
}
