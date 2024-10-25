package main

import (
	"fmt"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config, area string) error {
	fmt.Println()
	fmt.Println("location-area:")
	locations, err := pokeapi.PokeapiCall(cfg.next)
	if err != nil {
		fmt.Println()
		fmt.Println("can not get the locations")
		fmt.Println()
		return nil
	}
	cfg.next = *locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf(" -%s\n", location.Name)
	}
	fmt.Println()
	return nil
}
