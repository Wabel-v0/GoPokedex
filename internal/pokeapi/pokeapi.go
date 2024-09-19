package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mohalw44/gopokedex/internal/pokecache"
)

func (c *Client) LocationApi(pageUrl *string) (Pokedex, error) {
	var pokedex Pokedex
	url := baseURL + "location-area/"
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
