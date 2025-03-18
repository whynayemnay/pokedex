package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide pokemon name")
	}

	pokemonName := args[0]

	_, exists := cfg.pokeDex[pokemonName]
	if exists {
		fmt.Println("pokeman is already in you pokeDex: ", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...", pokemonName)
	pokemonEncounter, err := cfg.pokeapiClient.EncounterPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("...")
	time.Sleep(5 * time.Second)
	try := CatchPokemon(pokemonEncounter.BaseExperience)
	if try {
		fmt.Println(pokemonName, " was successfuly caught")
		cfg.pokeDex[pokemonName] = pokemonEncounter
		return nil
	}

	fmt.Println(pokemonName, " escaped")
	return nil
}

func CatchPokemon(baseExperience int) bool {
	catchChance := rand.Float64()

	// Convert baseExperience into a probability (higher baseExp = lower catch chance)
	catchProbability := 1.0 / (1.0 + float64(baseExperience)/100.0)

	return catchChance < catchProbability
}
