package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *Config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Prev

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("No previous page of locations")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Prev

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
