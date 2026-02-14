package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func commandCatch(cfg *Config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("Please provide a pokemon to catch")
	}

	pokemon := strings.ToLower(args[0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	catch := Catch{}
	parse := json.Unmarshal(body, &catch)
	if parse != nil {
		log.Fatal(parse)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	var caught bool
	var chance int
	var roll_count int
	var bonus int

	for {

		if catch.Experience <= 100 {
			chance += 70
		} else if catch.Experience <= 200 {
			chance += 50
		} else {
			chance += 30
		}

		switch roll_count {
		case 1:
			bonus = 5
		case 2:
			bonus = 10
		default:
			bonus = 0
		}

		roll := rand.Intn(306) + chance + bonus

		roll_count++

		if roll_count >= 3 {
			fmt.Printf("%s got away after 3 attempts!\n", catch.Name)
			break
		}

		fmt.Printf("Attempt %d: ", roll_count)

		if roll > catch.Experience {
			fmt.Printf("%s was caught!\n", catch.Name)
			caught = true
		} else if roll == catch.Experience {
			fmt.Printf("%s was caught, but it broke free!\n", catch.Name)
		} else {
			fmt.Printf("%s escaped!\n", catch.Name)
		}

		if caught {
			fmt.Printf("Congratulations! You caught %s with %d base experience!\n", catch.Name, catch.Experience)
			break
		}
	}

	if caught {
		captured[pokemon] = catch
	}

	return nil
}
