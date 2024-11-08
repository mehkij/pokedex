package main

import (
	"fmt"
	"os"

	"github.com/mehkij/pokedex/internal/pokeapi"
	"github.com/mehkij/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextAreaURL   *string
	prevAreaURL   *string
	pokeCache     pokecache.Cache
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    callbackMapb,
		},
	}
}

func callbackHelp(config *config, params []string) error {
	c := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, val := range c {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil

}

func callbackExit(config *config, params []string) error {
	os.Exit(0)
	return nil
}

func callbackMap(config *config, params []string) error {
	// Guards against empty config
	if config.nextAreaURL == nil && config.prevAreaURL != nil {
		fmt.Println("You are on the last page!")
		return nil
	}

	res, err := config.pokeapiClient.ListLocationAreas(config.pokeCache, config.nextAreaURL)

	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	config.nextAreaURL = res.Next
	config.prevAreaURL = res.Previous

	return nil
}

func callbackMapb(config *config, params []string) error {
	if config.prevAreaURL == nil {
		fmt.Println("You are on the first page!")
		return nil
	}

	res, err := config.pokeapiClient.ListLocationAreas(config.pokeCache, config.prevAreaURL)

	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	config.nextAreaURL = res.Next
	config.prevAreaURL = res.Previous

	return nil
}
