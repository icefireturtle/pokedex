package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func commandExplore(cfg *Config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("Please provide a location to explore")
	}

	location := strings.ToLower(args[0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Respons failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	explore := Explore{}
	parse := json.Unmarshal(body, &explore)
	if parse != nil {
		log.Fatal(parse)
	}

	fmt.Println("Found Pokemon:")

	for _, areas := range explore.PokemonEncounters {
		encounters := areas.Pokemon.Name
		fmt.Printf(" - %s\n", encounters)
	}

	return nil

}
