package main

import (
	"fmt"
	"strings"
)

var (
	sentences = [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars     = [5]rune{'o', 'E', 'S', 'П', 'М'}
)

func parseTest(sentences [4]string, chars [5]rune) (rt [4][5]int) {
	for i, sentence := range sentences {
		words := strings.Split(sentence, " ")
		for j := 0; j < 5; j++ {
			rt[i][j] = strings.Index(words[len(words)-1], string(chars[j]))
		}
	}

	return
}

func main() {
	foo := parseTest(sentences, chars)

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%v position %v\n", string(chars[j]), foo[i][j])
		}
		fmt.Println()
	}
}
