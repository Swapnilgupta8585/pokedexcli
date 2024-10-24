package main

import (
	"fmt"
	"math/rand"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, pokemonName string) error {

	baseUrl := "https://pokeapi.co/api/v2/pokemon"
	newUrl := baseUrl + "/" + pokemonName

	pokemon, err := pokeapi.PokeApiCall(newUrl)
	if err != nil {
		return fmt.Errorf("can not get the pokemon information: %w", err)
	}

	if _, ok := cfg.user[pokemonName]; ok {
		fmt.Printf("%s is already caught!\n", pokemonName)
		return nil
	}
	randomNumber := rand.Float64() * 20
	moreRandomNum := randomNumber + float64((pokemon.BaseExperience)/80)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if moreRandomNum <= 10.0 {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.user[pokemonName] = pokemon
		fmt.Println()
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		fmt.Println()
	}
	return nil

}
