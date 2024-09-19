package cmd

import (
	"fmt"
)

func commandExplore(cfg *Config, name string) error {
	fmt.Printf("You are now exploring the %s \n", name)

	return nil
}
