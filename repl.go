package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sharath070/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeApiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, ok := getCommands()[commandName]
		if ok {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
