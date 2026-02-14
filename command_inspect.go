package main

import (
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {

	if len(args[0]) == 0 {
		return fmt.Errorf("Please identify pokemon to inspect.")
	}

	if len(captured) == 0 {
		return fmt.Errorf("You haven't caught any pokemon yet!")
	}

	pokemon := args[0]

	if friend, exists := captured[pokemon]; exists {

		fmt.Println("You have caught this pokemon.")

		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", friend.Name, friend.Height, friend.Weight)

		fmt.Println("Stats:")

		for stat := range friend.Stats {
			fmt.Printf("-%s: %d\n", friend.Stats[stat].Stat.Name, friend.Stats[stat].BaseStat)
		}

		fmt.Println("Types:")

		for pokemon_type := range friend.Types {
			fmt.Printf("-%s\n", friend.Types[pokemon_type].Type.Name)
		}

	} else {

		fmt.Println("You have not caught this pokemon.")

	}

	return nil
}
