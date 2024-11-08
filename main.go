package main

import (
	"time"

	"github.com/mehkij/pokedex/internal/pokeapi"
	"github.com/mehkij/pokedex/internal/pokecache"
)

func main() {
	cnfig := config{
		pokeapiClient: pokeapi.NewClient(),
		pokeCache:     pokecache.NewCache(time.Second * 5),
	}

	startRepl(&cnfig)
}
