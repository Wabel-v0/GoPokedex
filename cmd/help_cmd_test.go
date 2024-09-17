package cmd

import "testing"

func TestCommandHelp(t *testing.T) {
	tests := []struct {
		name string
		usage string
		expected string
	}{
		{
			name: "help",
			usage: "Prints help message",
			expected: "help: Prints help message",
		},
		{
			name : "exit",
			usage: "Exits the program",
			expected: "exit: Exits the program",
		},
		{
			name : "list",
			usage: "Lists all the pokemons",
			expected: "list: Lists all the pokemons",
		},
		
		
	}	
	for i,tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			if tc.expected != tc.name + ": " + tc.usage {
				t.Errorf("Test %d failed: expected %s, got %s", i, tc.expected, tc.name + ": " + tc.usage)

			}
			
		})
	}
}