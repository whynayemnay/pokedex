package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	commands = map[string]cliCommnd{
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "display the help message",
			callback:    commandHelp,
		},
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
		command, exists := commands[word]
		if exists {
			err := command.callback()
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

type cliCommnd struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommnd
