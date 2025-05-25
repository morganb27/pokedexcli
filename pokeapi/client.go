package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/morganb27/pokedexcli/internal/config"
)

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
	}
}


func (c *Client) FetchLocations(pageURL *string) (LocationResponse, error) {
	url := config.BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return LocationResponse{}, fmt.Errorf("GET %s: bad status: %d", url, res.StatusCode)
	}

	var locationResponse LocationResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationResponse)
	if err != nil {
		return LocationResponse{}, err
	}

	return locationResponse, nil
}