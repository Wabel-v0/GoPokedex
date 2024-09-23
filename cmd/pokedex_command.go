package cmd

import "fmt"
func commandPokedex(cfg *Config, input string) error {
	for _, pokemon := range cfg.PokemonList {
		fmt.Println(pokemon.Name)
	}
	return nil
}