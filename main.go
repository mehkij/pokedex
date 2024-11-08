package main

import "github.com/mehkij/pokedex/internal/pokeapi"

func main() {
	cnfig := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cnfig)
}
