package cmd

import "fmt"

func commandHelp(cfg *Config) error {
	cmd := getCommands()
	for _, c := range cmd {
		fmt.Printf("%s: %s\n", c.name, c.usage)
	}

	return nil
}
