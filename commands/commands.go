package commands

import (
	"fmt"
	"os"

	"github.com/morganb27/pokedexcli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	PokeapiClient *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
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
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    mapfCommand,
		},
		"mapb": {
			name:        "map",
			description: "Get the previous page of locations",
			Callback:    mapbCommand,
		},
	}
}

func exitCommand(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	commands := GetCommands()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func mapfCommand(cfg *Config) error {
	locationsRes, err := cfg.PokeapiClient.FetchLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous
	fmt.Println("cfg.nextLocationsURL", cfg.nextLocationsURL)
	fmt.Println("cfg.prevLocationsURL", cfg.prevLocationsURL)
	fmt.Println("locationsRes.Results", locationsRes.Results)

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func mapbCommand(cfg *Config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	locationsRes, err := cfg.PokeapiClient.FetchLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous
	fmt.Println("cfg.nextLocationsURL", cfg.nextLocationsURL)
	fmt.Println("cfg.prevLocationsURL", cfg.prevLocationsURL)
	fmt.Println("locationsRes.Results", locationsRes.Results)

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

