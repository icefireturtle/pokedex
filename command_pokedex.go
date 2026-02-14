package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {

	if len(captured) == 0 {
		fmt.Println("You have not caught any pokemon!")
	} else {

		fmt.Println("Your Pokedex: ")

		for pokemon := range captured {
			fmt.Printf("- %s\n", pokemon)
		}
	}
	return nil
}
