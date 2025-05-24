package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Hello, World!")
}

func cleanInput(text string) []string  {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}