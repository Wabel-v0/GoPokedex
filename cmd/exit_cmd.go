package cmd

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Exiting the program.....")
	os.Exit(0)
	return nil
}
