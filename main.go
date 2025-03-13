package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for reader.Scan() {
		line := reader.Text()
		trim := cleanInput(line)
		word := trim[0]
		fmt.Printf("Your command was: %s\n", word)
		if err := reader.Err(); err != nil {
			fmt.Println("error reading input: ", err)
		}
		fmt.Print("Pokedex > ")

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
