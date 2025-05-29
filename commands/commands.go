package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/morganb27/pokedexcli/internal/pokecache"
	"github.com/morganb27/pokedexcli/pokeapi"
	"github.com/morganb27/pokedexcli/internal/config"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	PokeapiClient    *pokeapi.Client
	Cache            *pokecache.Cache
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
			name:        "mapb",
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
	fmt.Println("map() called")
	fmt.Printf("Before: prev = %v, next = %v\n", cfg.prevLocationsURL, cfg.nextLocationsURL)
	key := config.BaseURL + "/location-area"
	if cfg.nextLocationsURL != nil {
		key = *cfg.nextLocationsURL
	}

	cachedData, ok := cfg.Cache.Get(key)
	if ok {
		var res pokeapi.LocationResponse
		err := json.Unmarshal(cachedData, &res)
		if err != nil {
			return err
		}
		for _, location := range res.Results {
			fmt.Println(location.Name)
		}
		cfg.nextLocationsURL = res.Next
		cfg.prevLocationsURL = res.Previous
		fmt.Printf("After in CACHE: prev = %v, next = %v\n", cfg.prevLocationsURL, cfg.nextLocationsURL)
		return nil
	} 
	locationsRes, err := cfg.PokeapiClient.FetchLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	encoded, err := json.Marshal(locationsRes)
	if err != nil {
		return err
	}

	cfg.Cache.Add(key, encoded)
	fmt.Printf("After: prev = %v, next = %v\n", cfg.prevLocationsURL, cfg.nextLocationsURL)

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func mapbCommand(cfg *Config) error {
	fmt.Println("mapb() called")
	fmt.Printf("Before: prev = %v, next = %v\n", cfg.prevLocationsURL, cfg.nextLocationsURL)

	if cfg.prevLocationsURL == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	key := *cfg.prevLocationsURL
	cachedData, ok := cfg.Cache.Get(key)
	if ok {
		var res pokeapi.LocationResponse

		err := json.Unmarshal(cachedData, &res)
		if err != nil {
			return err
		}
		for _, location := range res.Results {
			fmt.Println(location.Name)
		}
		cfg.nextLocationsURL = res.Next
		cfg.prevLocationsURL = res.Previous
		fmt.Printf("After in CACHE: prev = %v, next = %v\n", cfg.prevLocationsURL, cfg.nextLocationsURL)
		return nil
	} 
	locationsRes, err := cfg.PokeapiClient.FetchLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous
	fmt.Printf("After: prev = %v, next = %v\n", cfg.prevLocationsURL, cfg.nextLocationsURL)


	encoded, err := json.Marshal(locationsRes)
	if err != nil {
		return err
	}

	cfg.Cache.Add(key, encoded)

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

