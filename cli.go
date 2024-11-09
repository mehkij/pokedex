package main

import (
	"fmt"
	"math/rand"
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
	pokedex       map[string]pokeapi.PokemonRes
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
		"explore": {
			name:        "explore",
			description: "Displays the available Pokemon in the given area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a given Pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon you have caught previously",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the Pokemon you have caught",
			callback:    callbackPokedex,
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
	if len(params) > 1 {
		fmt.Println("too many arguments")
		return nil
	}

	os.Exit(0)
	return nil
}

func callbackMap(config *config, params []string) error {
	if len(params) > 1 {
		fmt.Println("too many arguments")
		return nil
	}

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
		fmt.Println("- ", location.Name)
	}

	config.nextAreaURL = res.Next
	config.prevAreaURL = res.Previous

	return nil
}

func callbackMapb(config *config, params []string) error {
	if len(params) > 1 {
		fmt.Println("too many arguments")
		return nil
	}

	if config.prevAreaURL == nil {
		fmt.Println("You are on the first page!")
		return nil
	}

	res, err := config.pokeapiClient.ListLocationAreas(config.pokeCache, config.prevAreaURL)

	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println("- ", location.Name)
	}

	config.nextAreaURL = res.Next
	config.prevAreaURL = res.Previous

	return nil
}

func callbackExplore(config *config, params []string) error {
	// If len == 1, user only passed command name with no arguments
	if len(params) == 1 {
		fmt.Println("not enough arguments")
		return nil
	}

	// params[0] is always the name of the command
	URL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", params[1])

	res, err := config.pokeapiClient.ListPokemonEncounters(config.pokeCache, &URL)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(res.PokemonEncounters) == 0 {
		fmt.Println("no Pokemon in this area")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Println("- ", encounter.Pokemon.Name)
	}

	return nil
}

func callbackCatch(config *config, params []string) error {
	if len(params) == 1 {
		fmt.Println("not enough arguments")
		return nil
	}

	_, ok := config.pokedex[params[1]]

	if ok {
		fmt.Printf("You already have %s!\n", params[1])
		return nil
	}

	// params[0] is always the name of the command
	URL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", params[1])

	res, err := config.pokeapiClient.GetPokemon(config.pokeCache, &URL)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", res.Name)

	if rand.Float64()*float64(res.BaseExperience) > float64(res.BaseExperience)/2 {
		fmt.Printf("You caught %v!\n", res.Name)
		config.pokedex[res.Name] = res
	} else {
		fmt.Printf("Drats! The %v got away!\n", res.Name)
	}

	return nil
}

func callbackInspect(config *config, params []string) error {
	if len(params) == 1 {
		fmt.Println("not enough arguments")
		return nil
	}

	pokemon, ok := config.pokedex[params[1]]

	if ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("- %s\n", t.Type.Name)
		}
	} else {
		fmt.Println("You have not caught this Pokemon yet!")
	}

	return nil
}

func callbackPokedex(config *config, params []string) error {
	if len(params) > 1 {
		fmt.Println("too many arguments")
		return nil
	}

	if len(config.pokedex) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	for _, pokemon := range config.pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
