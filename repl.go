package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
			err:= command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	clean := strings.ToLower(strings.TrimSpace(text))
	words := strings.Fields(clean)
	return words
}
	
type cliCommand struct {
	name string
	description string
	callback func(cfg *Config) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of the next 20 location areas in the Pokemon World",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon World",
			callback: commandMapBack,
		},
	}
}

type Locations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Config struct {
	Next string
	Previous string
}