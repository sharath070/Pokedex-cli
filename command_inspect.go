package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemon name has not been provided")
	}

	name := args[0]
	data, ok := cfg.caughtPokemon[name]
	if ok {
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Height: %d\n", data.Height)
		fmt.Printf("Weight: %d\n", data.Weight)
		fmt.Println("Stats:")
		for _, stat := range data.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, ability := range data.Types {
			fmt.Printf(" - %s\n", ability.Type.Name)
		}

	} else {
		fmt.Println("You have not caught that pokemon")
	}

	return nil
}
