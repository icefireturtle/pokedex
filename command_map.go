package main

import ( 
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)
func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != "" {
		url = cfg.Next
	}

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
	
	locations := Locations{}
	parse := json.Unmarshal(body, &locations)
	if parse != nil {
		log.Fatal(parse)
	}

	if locations.Next != nil {
		cfg.Next = *locations.Next
	} else {
		cfg.Next = ""
	}

	if locations.Previous != nil {
		cfg.Previous = *locations.Previous
	} else {
		cfg.Previous = ""
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

