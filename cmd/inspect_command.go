package cmd

import (
	"fmt"
)

func commandInspect(c *Config, userInput string) error {
	_, ok := c.PokemonList[userInput]
	if !ok {
		fmt.Println("pokemon not found")
		return nil
	}
	for _, pokemon := range c.PokemonList {
		fmt.Printf("name : %v \n", pokemon.Name)
		fmt.Printf("Hight: %v \n", pokemon.Height)
		fmt.Printf("Weight: %v \n", pokemon.Weight)
	}
	return nil
}
