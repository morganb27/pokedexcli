package main

import (
	"github.com/morganb27/pokedexcli/repl"
	"github.com/morganb27/pokedexcli/pokeapi"
	"github.com/morganb27/pokedexcli/commands"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
	}
    repl.StartREPL(cfg)
}