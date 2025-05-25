package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    helpCommand,
		},
	}
}

func exitCommand() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	commands := GetCommands()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}