package main

import "fmt"

func main() {
	var (
		first, second int
	)

	fmt.Println("Enter two values - first must be less than second!")
	fmt.Println("Enter first value")
	fmt.Scan(&first)

	fmt.Println("Enter second value")
	fmt.Scan(&second)

	if first >= second {
		fmt.Println("I said first must be less than second!")
		return
	}

	sum := first + second

	for i := first; i < sum+1; i++ {
		fmt.Printf("Current step is - %v\n", i)
	}
}
