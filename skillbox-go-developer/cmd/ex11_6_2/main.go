package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	message = "a10 10 20b 20 30c30 30 dd"
)

func main() {
	var (
		words   []string
		numbers []int64
	)

	words = strings.Split(message, " ")

	for _, word := range words {
		if i, err := strconv.ParseInt(word, 10, 64); err == nil {
			numbers = append(numbers, i)
		}
	}

	fmt.Println("String contains a numbers -", numbers)
}
