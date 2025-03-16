package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/whynayemnay/pokedex/internal/pokeapi"
	"github.com/whynayemnay/pokedex/internal/pokecache"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	cache := pokecache.NewCache(60 * time.Second)
	pokedexClient := pokeapi.NewClient(5*time.Second, cache)
	cfg := &config{
		pokeapiClient: pokedexClient,
		cache:         cache,
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
		arg := line[1:]

		command, exists := getCommands()[word]
		if exists {
			err := command.callback(cfg, arg)
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
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient *pokeapi.Client
	cache         *pokecache.Cache
	nextUrl       *string
	previousUrl   *string
}

// var configGlobal = &config{}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(cfg *config, args []string) error { return commandHelp(cfg) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(cfg *config, args []string) error { return commandExit(cfg) },
		},
		"map": {
			name:        "map",
			description: "Show 20 locations",
			callback:    func(cfg *config, args []string) error { return commandMapF(cfg) },
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 locations",
			callback:    func(cfg *config, args []string) error { return commandMapB(cfg) },
		},
		"explore": {
			name:        "explore",
			description: "Explore the given location, showing Pokémon found there",
			callback:    commandExplore, // Now correctly accepts args
		},
	}
}
