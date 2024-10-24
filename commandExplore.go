package main

import (
	"fmt"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config, area string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	newUrl := baseUrl + "/" + area

	pokemons, err := pokeapi.ExploreApiCall(newUrl)
	if err != nil {
		return fmt.Errorf("can not get the locations: %w", err)
	}
	fmt.Println()
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	fmt.Println()
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil

}
