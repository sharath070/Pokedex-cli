package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	fmt.Printf("Exploring %s...\n", name)

	locationRes, err := cfg.pokeApiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Println("Found pokemon:")
	for _, encounter := range locationRes.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
