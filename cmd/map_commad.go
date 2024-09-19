package cmd

import "fmt"

func commandMap(cfg *Config, name string) error {
	resp, err := cfg.Client.LocationApi(cfg.NextURL)
	if err != nil {
		return err
	}
	cfg.NextURL = resp.Next
	cfg.PrevURL = resp.Previous
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *Config, name string) error {
	if cfg.PrevURL == nil {
		fmt.Println("No previous page")
		return nil
	}
	resp, err := cfg.Client.LocationApi(cfg.PrevURL)
	if err != nil {
		return err
	}
	cfg.NextURL = resp.Next
	cfg.PrevURL = resp.Previous
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
