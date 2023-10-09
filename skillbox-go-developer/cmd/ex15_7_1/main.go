package main

import "fmt"

var (
	arr = [...]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}
)

func isEven(a int) bool {
	if (a % 2) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var (
		even, odd int
	)

	for _, i := range arr {
		if isEven(i) {
			even++
		} else {
			odd++
		}
	}

	fmt.Printf("even - %v, odd - %v\n", even, odd)
}
