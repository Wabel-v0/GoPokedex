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
	callback func(*Config, string) error
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
		"explore": {
			name:     "explore",
			usage:    "Explore the pokemon world",
			callback: commandExplore,
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

		cmd, ok := getCommands()[userInput[0]]
		// checing the user input
		if len(userInput) == 0 {
			continue
		}
		if cmd.name == "explore" && len(userInput) == 1 {
			fmt.Println("Please provide a name to explore")
			continue
		}
		if cmd.name == "explore" && len(userInput) > 2 {
			fmt.Println("Make sure there are no spaces in the name")
			continue
		}
		// to be continued
		if ok {
			err := cmd.callback(cfg, userInput[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
