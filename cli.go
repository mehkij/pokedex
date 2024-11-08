package main

import (
	"fmt"
	"os"

	"github.com/mehkij/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextAreaURL   *string
	prevAreaURL   *string
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

func callbackHelp(config *config) error {
	c := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, val := range c {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil

}

func callbackExit(config *config) error {
	os.Exit(0)
	return nil
}

func callbackMap(config *config) error {
	if config.nextAreaURL == nil {
		fmt.Println("You are on the last page!")
		return nil
	}

	res, err := config.pokeapiClient.ListLocationAreas(config.nextAreaURL)

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

func callbackMapb(config *config) error {
	if config.prevAreaURL == nil {
		fmt.Println("You are on the first page!")
		return nil
	}

	res, err := config.pokeapiClient.ListLocationAreas(config.prevAreaURL)

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
