package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	return words
}
