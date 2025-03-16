package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide the name of the pokemon")
	}

	name := args[0]
	pokemon, err := cfg.pokeApiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	curLevel := rand.Intn(256)
	if pokemon.BaseExperience > curLevel {
		fmt.Printf("%s esacaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[name] = pokemon
	}

	return nil
}
