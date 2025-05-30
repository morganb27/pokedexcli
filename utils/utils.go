package utils

import (
	"strings"
)

func CleanInput(text string) []string  {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}