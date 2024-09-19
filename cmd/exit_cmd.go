package cmd

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, name string) error {
	fmt.Println("Exiting the program.....")
	os.Exit(0)
	return nil
}
