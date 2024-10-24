package main

import (
	"github.com/Swapnilgupta8585/pokedexcli/internal/pokeapi"
)

func main() {
	var cfg = &Config{
		next:     "https://pokeapi.co/api/v2/location-area",
		previous: nil,
		user:     make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)

}
