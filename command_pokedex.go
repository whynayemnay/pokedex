package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokeDex) == 0 {
		return errors.New("you havent cought any pokemon yet, your pokedex is empty")
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.pokeDex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
