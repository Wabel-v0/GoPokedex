package cmd

import (
	"fmt"
)

func commandExplore(cfg *Config, name string) error {
	fmt.Printf("You are now exploring the %s \n", name)
	explore, err := cfg.Client.PokemonEncountersApi(name)
	if err != nil {
		return err
	}
	for _, poke := range explore.PokemonEncounters {
		fmt.Printf("Pokemon Name: %s \n", poke.Pokemon.Name)
	}
	return nil
}
