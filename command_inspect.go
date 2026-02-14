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

	for owned := range captured {

		if owned == pokemon {
			c := captured[owned]
			fmt.Println("You have caught this pokemon.")

			fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", c.Name, c.Height, c.Weight)

			fmt.Println("Stats:")

			for stat := range c.Stats {
				fmt.Printf("-%s: %d\n", c.Stats[stat].Stat.Name, c.Stats[stat].BaseStat)
			}

			fmt.Println("Types:")

			for ptype := range c.Types {
				fmt.Printf("-%s\n", c.Types[ptype].Type.Name)
			}
		} else {
			return fmt.Errorf("You have not caught this pokemon.")
		}
	}
	return nil
}
