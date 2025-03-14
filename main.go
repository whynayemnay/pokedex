package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

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
		command, exists := getCommands()[word]
		if exists {
			err := command.callback(configGlobal)
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
	callback    func(*config) error
}

type config struct {
	nextUrl     string
	previousUrl string
}

var configGlobal = &config{}

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
			description: "show 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show previous 20 locations",
			callback:    commandMapB,
		},
	}
}
