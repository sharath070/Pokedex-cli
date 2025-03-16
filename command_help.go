package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	commands := getCommands()

	fmt.Printf("Welcome to the Pokedex!\n\n")
	fmt.Print("Useage\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
