package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	message = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
)

func main() {
	var (
		words     []string
		countdown int
	)

	words = strings.Split(message, " ")

	for _, word := range words {
		if unicode.IsUpper(rune(strings.Trim(word, " ,")[0])) {
			countdown++
		}
	}

	fmt.Printf("Mesage contains a - %v words starts with capital letters\n", countdown)
}
