package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfoApi(pathURL string) (PokemonInfo, error) {
	url := baseURL + "pokemon/" + pathURL

	resp, err := http.Get(url)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	var pokemonInfo PokemonInfo
	if err = json.Unmarshal(data, &pokemonInfo); err != nil {
		return PokemonInfo{}, err

	}
	return pokemonInfo, nil
}
