package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide location name")
	}

	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)

	location, err := cfg.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}

	// Extract Pokémon names
	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found in this location.")
		return nil
	}

	fmt.Println("Found Pokémon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
