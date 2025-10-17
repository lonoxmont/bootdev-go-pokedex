package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},

		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch for input %q: got %d, want %d; actual=%v expected=%v",
				c.input, len(actual), len(c.expected), actual, c.expected)
			continue // avoid out-of-range index
		}
		for i := range actual {

			if actual[i] != c.expected[i] {
				t.Errorf("word mismatch at %d for input %q: wanted %q got %q", i, c.input, actual[i], c.expected[i])
			}
		}
	}
}
