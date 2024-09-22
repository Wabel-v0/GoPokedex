package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mohalw44/gopokedex/internal/pokecache"
)

func (c *Client) PokemonEncountersApi(pathURL string) (PokemonEncounters, error) {
	url := baseURL + "location-area/" + pathURL
	var pokemonEncounters PokemonEncounters
	cashe := pokecache.NewCache(5)
	getCashe, ok := cashe.Get(url)
	if ok {
		json.Unmarshal(getCashe, &pokemonEncounters)
		return pokemonEncounters, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return PokemonEncounters{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, err
	}
	cashe.Add(url, data)
	if err = json.Unmarshal(data, &pokemonEncounters); err != nil {
		return PokemonEncounters{}, err
	}
	return pokemonEncounters, nil
}
