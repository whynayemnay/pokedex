package main

import "testing"

func TestCleanupInput(t *testing.T) {
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
			input:    "   pikachu    ",
			expected: []string{"pikachu"},
		},
		{
			input:    "   bulbasaur   charmander   ",
			expected: []string{"bulbasaur", "charmander"},
		},
		{
			input:    "apple banana    orange",
			expected: []string{"apple", "banana", "orange"},
		},
		{
			input:    "test   input",
			expected: []string{"test", "input"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		t.Logf("input: '%s', expected '%v'", c.input, c.expected)

		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("For input '%s', expected word '%s' but got '%s'", c.input, c.expected[i], actual[i])
			}
		}
	}
}
