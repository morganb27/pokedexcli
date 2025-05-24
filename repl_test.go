package main

import (
	"testing"
	"fmt"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic split",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "split with uppercase letters",
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
            t.Errorf("FAIL [%s]: input=%q expected=%v got=%v", c.name, c.input, c.expected, actual)
        } else {
            fmt.Printf("PASS [%s]\n", c.name)
        }
	}
    
}