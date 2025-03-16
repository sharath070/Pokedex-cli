package main

import (
	"time"

	"github.com/sharath070/pokedex-cli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := config{
		pokeApiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
