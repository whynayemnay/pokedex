package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = locations.Next
	cfg.previousUrl = locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousUrl == nil {
		return errors.New("you are on the first page")
	}
	locations, err := cfg.pokeapiClient.ListLocations(cfg.previousUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = locations.Next
	cfg.previousUrl = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
