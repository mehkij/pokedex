package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mehkij/pokedex/internal/pokecache"
)

func (c *Client) ListLocationAreas(cache pokecache.Cache, URL *string) (LocationAreasRes, error) {
	var fullURL string
	endpoint := "/location-area"

	if URL != nil {
		fullURL = *URL
	} else {
		fullURL = baseURL + endpoint
	}

	val, ok := cache.Get(fullURL)

	if ok {
		var cachedResponse LocationAreasRes

		e := json.Unmarshal(val, &cachedResponse)

		return cachedResponse, e
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasRes{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasRes{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreasRes{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasRes{}, err
	}

	var response LocationAreasRes
	e := json.Unmarshal(data, &response)

	cache.Add(fullURL, data)

	return response, e
}
