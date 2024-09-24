package cmd

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *Config, input string) error {
	fmt.Println("Catching a pokemon")
	poke, err := cfg.Client.PokemonInfoApi(input)
	if err != nil {
		return err
	}

	chance := rand.Intn(poke.BaseExperience)

	time.Sleep(1 * time.Second)
	if chance > 35 {
		fmt.Printf("%s ran away \n", poke.Name)
		return nil

	}
	fmt.Printf("You caught %s \n", poke.Name)
	cfg.PokemonList[poke.Name] = poke

	return nil

}
