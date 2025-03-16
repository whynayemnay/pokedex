package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: explore <location-area>")
		return nil
	}

	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)

	location, err := cfg.pokeapiClient.ListPokemon(&locationName)
	if err != nil {
		return fmt.Errorf("could not explore location: %w", err)
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
