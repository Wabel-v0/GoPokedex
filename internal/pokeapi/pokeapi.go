package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mohalw44/gopokedex/internal/pokecache"
)

func (c *Client) CallPokeApi(pageUrl *string) (Pokedex, error) {
	var pokedex Pokedex
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}
	cache := pokecache.NewCache(5)

	getCache, ok := cache.Get(url)
	if ok {
		json.Unmarshal(getCache, &pokedex)
		return pokedex, nil
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
	cache.Add(url, data)
	if err = json.Unmarshal(data, &pokedex); err != nil {
		return Pokedex{}, err
	}
	return pokedex, nil

}
