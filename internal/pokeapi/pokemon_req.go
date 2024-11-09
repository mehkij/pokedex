package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mehkij/pokedex/internal/pokecache"
)

func (c *Client) GetPokemon(cache pokecache.Cache, URL *string) (PokemonRes, error) {
	val, ok := cache.Get(*URL)

	// If the data already exists in the cache, use it instead of fetch
	if ok {
		var cachedResponse PokemonRes

		e := json.Unmarshal(val, &cachedResponse)

		return cachedResponse, e
	}

	req, err := http.NewRequest("GET", *URL, nil)

	if err != nil {
		return PokemonRes{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return PokemonRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonRes{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonRes{}, err
	}

	var response PokemonRes
	e := json.Unmarshal(data, &response)

	cache.Add(*URL, data)

	return response, e
}
