# pokedex

A CLI Pokedex tool using PokeAPI built entirely in Go.

## Usage

Using the help command will return a list of all available commands and their descriptions.

```
help
```

You can also type the name of a command after to get the description of that specific command:

```
help catch
Usage:
catch: Attempt to catch a given Pokemon
```
## Changelog

### v1.0.0

The release version! This version comes with a basic set of commands.

Commands that fetch from the API are cached. The cache is cleared every 5 seconds by default, but this can be changed by passing in a new time.Duration argument to the pokeCache field in the config in the main.go file:

```go
func main() {
	cnfig := config{
		pokeapiClient: pokeapi.NewClient(),
		pokeCache:     pokecache.NewCache(time.Second * 5),
		pokedex:       map[string]pokeapi.PokemonRes{},
	}

	// ...
}
```
