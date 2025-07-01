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
			input:    "Hello world!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    "		  hello world \n ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "CATS RULE!!! ",
			expected: []string{"cats", "rule!!!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Expected length %d does not match actual length %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word '%s' does not match actual word '%s'", expectedWord, word)
			}
		}
	}
}
