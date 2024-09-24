package cmd

import "fmt"

func commandPokedex(cfg *Config, input string) error {
	fmt.Println("Pokedex:")
	for _, pokemon := range cfg.PokemonList {
		fmt.Printf("- %s \n", pokemon.Name)
	}
	return nil
}
