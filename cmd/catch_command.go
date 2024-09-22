package cmd

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, input string) error {
	fmt.Println("Catching a pokemon")
	poke, err := cfg.Client.PokemonInfoApi(input)
	if err != nil {
		return err
	}
	r := rand.Float64()
	chance := min(0.9, float64(poke.BaseExperience)/100)
	fmt.Println(r, chance)
	if r > chance {
		fmt.Println("You caught a", poke.Name)

	} else {
		fmt.Println("The pokemon ran away")
	}

	return nil

}
