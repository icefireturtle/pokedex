package main

import "fmt"

func commandHelp(cfg *Config, args ...string) error {
	//Header
	fmt.Println()
	fmt.Println("Pokedex CLI Help")
	fmt.Println("Usage:")
	fmt.Println()

	for _, directive := range commands {
		fmt.Printf("%s: %s\n", directive.name, directive.description)
	}
	return nil
}
