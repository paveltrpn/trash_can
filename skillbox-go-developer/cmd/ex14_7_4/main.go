package main

import "fmt"

var (
	varOne   float64 = 10.1
	varTwo   float64 = 20.2
	varThree float64 = 30.3
)

func foo(a float64) float64 {
	return a + varOne
}

func bar(a float64) float64 {
	return a + varTwo
}

func baz(a float64) float64 {
	return a + varThree
}

func main() {
	var (
		input float64
	)

	fmt.Println("Enter a number")
	fmt.Scan(&input)

	output := foo(bar(baz(input)))

	fmt.Printf("output is - %v\n", output)
}
