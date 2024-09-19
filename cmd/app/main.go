package main

import (
	"time"

	"github.com/mohalw44/gopokedex/cmd"
	"github.com/mohalw44/gopokedex/internal/pokeapi"
)

func main() {
	newClient := pokeapi.NewClient(5 * time.Second)
	cfg := &cmd.Config{
		Client: newClient,
	}
	cmd.Run(cfg)

}
