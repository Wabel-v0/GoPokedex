package cmd

import (
	"testing"
)
func TestCliCommand(t *testing.T) {
	tests := []struct {
		input string
		expected []string
	}{
		{
			input: "help",
			expected: []string{"help"},
		},
		{
			input: "    ",
			expected: []string{},
		},
		{
			input: "baba json",
			expected: []string{"baba", "json"},
		},
		{
			input: "exit from the command",
			expected: []string{"exit", "from", "the", "command"},
		},
	}
	for i,tc := range tests {
		t.Run("TestCliCommand", func(t *testing.T) {
			got := cleanInput(tc.input)
			if len(got) != len(tc.expected) {
				t.Errorf("Test %d: Expected %v, got %v", i, tc.expected, got)
			}
			for i, word := range got {
				if word != tc.expected[i] {
					t.Errorf("Test %d: Expected %v, got %v", i, tc.expected, got)
				}
			}
		

	})

}
}