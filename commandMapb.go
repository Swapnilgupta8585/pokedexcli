package main

import (
	"fmt"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *Config, area string) error {
	if cfg.previous != nil {
		locations, err := pokeapi.PokeapiCall(*cfg.previous)
		if err != nil {
			fmt.Println()
			fmt.Println("can not get the locations")
			fmt.Println()
			return nil
		}
		fmt.Println()
		fmt.Println("location-area:")
		cfg.next = *locations.Next
		cfg.previous = locations.Previous
		for _, location := range locations.Results {
			fmt.Printf(" -%s\n", location.Name)
		}
		fmt.Println()
		return nil
	}
	fmt.Println()
	fmt.Println("can not view previous page as it has no locations")
	fmt.Println()
	return nil
}
