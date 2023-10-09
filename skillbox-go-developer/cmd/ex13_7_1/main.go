package main

import "fmt"

func isEven(a int) bool {
	if (a % 2) == 0 {
		return true
	} else {
		return false
	}
}

func foo(a, b int) {
	var (
		accum int
	)

	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		if (i % 2) == 0 {
			accum += i
		}
	}

	fmt.Println(accum)
}

func main() {
	foo(4, 1)
}
