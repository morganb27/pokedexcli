package main

import (
	"time"

	"github.com/morganb27/pokedexcli/commands"
	"github.com/morganb27/pokedexcli/internal/pokecache"
	"github.com/morganb27/pokedexcli/pokeapi"
	"github.com/morganb27/pokedexcli/repl"
)

func main() {
	pokeClient := pokeapi.NewClient()
	pokeCache := pokecache.NewCache(30 * time.Second)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
		Cache: pokeCache,
	}
    repl.StartREPL(cfg)
}