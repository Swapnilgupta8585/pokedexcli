package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	next     string
	previous *string
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Let's you know the next 20 locations to explore",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Let's you know the previous 20 locations to explore",
			callback:    commandMapb,
		},
	}
}
