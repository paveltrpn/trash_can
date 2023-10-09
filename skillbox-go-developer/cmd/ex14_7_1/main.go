package main

import "fmt"

func isEven(a int) bool {
	if (a % 2) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var (
		foo int = 10
	)

	fmt.Printf("is even - %v\n", isEven(foo))
}
