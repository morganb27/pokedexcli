package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"

	"github.com/morganb27/pokedexcli/commands"
)

func main() {
    startREPL()
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	commands := commands.GetCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan() 
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		cmd, ok := commands[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.Callback(); err != nil {
			fmt.Printf("Error with command %s: %v", commandName, err)
		}
	}
}

func cleanInput(text string) []string  {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}