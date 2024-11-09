package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mehkij/pokedex/internal/pokecache"
)

func (c *Client) ListPokemonEncounters(cache pokecache.Cache, URL *string) (AreaEncounterRes, error) {
	val, ok := cache.Get(*URL)

	if ok {
		var cachedResponse AreaEncounterRes

		e := json.Unmarshal(val, &cachedResponse)

		return cachedResponse, e
	}

	req, err := http.NewRequest("GET", *URL, nil)

	if err != nil {
		return AreaEncounterRes{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return AreaEncounterRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return AreaEncounterRes{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return AreaEncounterRes{}, err
	}

	var response AreaEncounterRes
	e := json.Unmarshal(data, &response)

	cache.Add(*URL, data)

	return response, e
}
