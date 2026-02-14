package main

import (
	"fmt"
	"time"

	"github.com/icefireturtle/pokedex/internal/pokeapi"
)

func main() {
	fmt.Println("Welcome to the Pokedex!")
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	startREPL(cfg)
}
