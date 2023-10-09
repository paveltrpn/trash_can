package main

import "fmt"

func main() {
	var (
		value int
	)

	fmt.Println("Enter positive integer value")
	fmt.Scan(&value)

	if value <= 0 {
		fmt.Println("I said a positive value!")
		return
	}

	for i := 0; i < value; i++ {
		fmt.Printf("iteration number is - %v\n", i)
	}
}
