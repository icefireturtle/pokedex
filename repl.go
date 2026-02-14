package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	cfg := &Config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean := cleanInput(scanner.Text())
		if len(clean) == 0 {
			continue
		}

		command, exists := commands[clean[0]]
		if exists {
			err := command.callback(cfg, clean[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	clean := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(clean)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon World",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explores the map and finds pokeman in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the pokemon encountered",
			callback:    commandCatch,
		},
	}
}

type Locations struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Next     string
	Previous string
}

type Explore struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Catch struct {
	Name       string `json:"name"`
	Experience int    `json:"base_experience"`
}

var captured map[string]Catch = make(map[string]Catch)
