
# 🐦‍🔥 pokedexcli

>  A sleek and interactive CLI tool to explore the Pokémon world — catch, inspect, and navigate your way through the Pokédex right from your terminal. Built with Go.

## ✨ Features

- 🔍 **Explore Pokémon** regions with intuitive CLI commands  
- 🗺️ **Map Navigation**: Move through different Pokémon locations  
- 🎒 **Pokédex**: Catch and inspect Pokémon as you encounter them  
- 🧠 **REPL Interface**: Seamlessly interact with commands in a shell-like environment  
- 💾 **In-Memory Caching**: Reduces API calls by caching responses for 60 seconds 
- 📦 **Modular Design**: Clean, scalable Go architecture


## 🧰 Commands

| Command            | Description                          |
|--------------------|--------------------------------------|
| `help`             | Show available commands              |
| `map`              | Show current map and directions      |
| `mapb`             | Move back in the map                 |
| `explore <area-name>`| Explore and discover Pokémon         |
| `catch <name>`     | Catch a Pokémon                      |
| `inspect <name>`   | Inspect a caught Pokémon             |
| `pokedex`          | Show all caught Pokémon              |
| `exit`             | Quit the application                 |

## 🚀 Getting Started

### Prerequisites

- [Go 1.23.2](https://go.dev/dl/) or higher

## Installation

```bash
git clone https://github.com/Swapnilgupta8585/pokedexcli
cd pokedexcli
go run .
```

## 🏗️ Project Structure
```
pokedexcli/
├── main.go # App entry point & REPL launch
├── repl.go # REPL loop and command parsing
├── commands.go # Command registration and lookup
├── command*.go # Individual command implementations
│ ├── commandCatch.go
│ ├── commandExplore.go
│ ├── commandInspect.go
│ ├── commandMap.go
│ └── ...etc
├── internal/
│ └── pokecache/
│          └── cache.go # In-memory cache implementation
│ └── pokeapi/
│          └── apiCall.go # General-purpose API handler with cache
│          └── apiExploreCall.go # Explore endpoint handler with caching
│          └── apiPokemonCall.go # Pokémon data fetching with caching
├── go.mod # Go module metadata
└── README.md # Project documentation
```

## 🧑‍💻 Contributing

Contributions are welcome!  
Feel free to open issues, submit PRs, or suggest new features.

