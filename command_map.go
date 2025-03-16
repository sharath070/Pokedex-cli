package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationResp.Next
	cfg.prevLocationUrl = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
