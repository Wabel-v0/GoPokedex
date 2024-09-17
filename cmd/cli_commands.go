package cmd

import (
	"bufio"
	"fmt"
	"os"
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

func Run(){
	
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		userInput := scanner.Text()
		for _, cmd := range commands {
			if userInput == cmd.name {
				err := cmd.callback()
				if err != nil {
					fmt.Println("Error:", err)
					
				}
			}
			
		}
	}
}
