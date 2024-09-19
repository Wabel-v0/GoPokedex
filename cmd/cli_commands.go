package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mohalw44/gopokedex/internal/pokeapi"
)

type Config struct {
	Client  pokeapi.Client
	NextURL *string
	PrevURL *string
}

type CliCommand struct {
	name     string
	usage    string
	callback func(*Config) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:     "help",
			usage:    "Prints help message",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			usage:    "Exits the program",
			callback: commandExit,
		},
		"map": {
			name:     "map",
			usage:    "Prints the map of the pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			usage:    "Prints the previous map of the pokemon world",
			callback: commandMapb,
		},
	}

}

func cleanInput(input string) []string {
	word := strings.ToLower(input)
	words := strings.Fields(word)

	return words

}

func Run(cfg *Config) {
	fmt.Println("Welcome to the Pokedex CLI")
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}
		cmd, ok := getCommands()[userInput[0]]
		if ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
