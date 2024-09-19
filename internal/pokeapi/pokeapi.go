package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CallPokeApi(pageUrl *string) (Pokedex, error) {
	{
		url := "https://pokeapi.co/api/v2/location-area/"
		if pageUrl != nil {
			url = *pageUrl
		}
		resp, err := http.Get(url)
		if err != nil {
			return Pokedex{}, err

		}
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokedex{}, err

		}
		var pokedex Pokedex
		if err = json.Unmarshal(data, &pokedex); err != nil {
			return Pokedex{}, err
		}
		return pokedex, nil
	}
}
