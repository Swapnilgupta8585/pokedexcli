package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, pokemonName string) error {

	baseUrl := "https://pokeapi.co/api/v2/pokemon"
	newUrl := baseUrl + "/" + pokemonName

	pokemon, err := pokeapi.PokeApiCall(newUrl)
	if err != nil {
		fmt.Println()
		fmt.Println("can not get the pokemon information")
		fmt.Println()
		return nil
	}
	fmt.Println()
	if _, ok := cfg.user[pokemonName]; ok {
		fmt.Printf("%s is already caught!\n", pokemonName)
		fmt.Println()
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	time.Sleep(1 * time.Second)
	randomNumber := rand.Float64() * 20
	moreRandomNum := randomNumber + float64((pokemon.BaseExperience)/80)
	if moreRandomNum <= 10.0 {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.user[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
		fmt.Println()
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		fmt.Println()
	}
	return nil

}
