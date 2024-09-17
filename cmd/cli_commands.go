package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type CliCommand struct {
	name string
	usage string
	callback func() error
}



func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name: "help",
			usage: "Prints help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			usage: "Exits the program",
			callback: commandExit,
		},
	}
	
}

func cleanInput(input string) []string {
	word := strings.ToLower(input)
	words := strings.Fields(word)
	
	return words

}

func Run(){
	fmt.Println("Welcome to the Pokedex CLI")
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}
		cmd , ok := getCommands()[userInput[0]]
		if ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		
		
}
}
