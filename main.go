package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/whynayemnay/pokedex/internal/pokeapi"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	pokedexClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokedexClient,
		pokeDex:       make(map[string]pokeapi.RespPokemonEncounter),
	}

	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			break
		}

		line := cleanInput(reader.Text())
		if len(line) == 0 {
			continue
		}
		word := line[0]
		args := []string{}
		if len(line) > 1 {
			args = line[1:]
		}

		command, exists := getCommands()[word]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("error", err)
			}
		} else {
			fmt.Println("unknown command")
		}

		if err := reader.Err(); err != nil {
			fmt.Println("error reading input: ", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextUrl       *string
	previousUrl   *string
	pokeDex       map[string]pokeapi.RespPokemonEncounter
}

// var configGlobal = &config{}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show 20 locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore the given location, showing pokemon found there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch the given pokemon",
			callback:    commandCatch,
		},
	}
}
