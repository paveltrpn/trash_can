package main

import "fmt"

func swap(a, b *int) {
	(*a), (*b) = (*b), (*a)
}

func main() {
	var (
		a, b int
	)

	a = 10
	b = 5

	fmt.Printf("before swap a = %v, b = %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap a = %v, b = %v\n", a, b)
}
