package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasRes, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

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

	return response, e
}
