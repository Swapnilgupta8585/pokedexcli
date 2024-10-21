package main

import (
	"fmt"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *Config) error {
	if cfg.previous != nil {
		locations, err := pokeapi.PokeapiCall(*cfg.previous)
		if err != nil {
			return fmt.Errorf("can not get the locations: %w", err)
		}
		fmt.Println()
		cfg.next = *locations.Next
		cfg.previous = locations.Previous
		for _, location := range locations.Results {
			fmt.Printf("Pokedex > %s\n", location.Name)
		}
		fmt.Println()
		return nil
	}
	return fmt.Errorf("can not view previous page as it has no locations")
}
