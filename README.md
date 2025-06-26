
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
| `explore`          | Explore and discover Pokémon         |
| `catch <name>`     | Catch a Pokémon                      |
| `inspect <name>`   | Inspect a caught Pokémon             |
| `pokedex`          | Show all caught Pokémon              |
| `exit`             | Quit the application                 |

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

### Installation

```bash
git clone https://github.com/Swapnilgupta8585/pokedexcli
cd pokedexcli
go run main.go
```

## 🏗️ Project Structure
```
├── command*.go        # Individual CLI command implementations
├── repl.go            # REPL environment logic
├── main.go            # App entry point
├── go.mod             # Go module file
└── README.md
```

## 🧑‍💻 Contributing

Contributions are welcome!  
Feel free to open issues, submit PRs, or suggest new features.

