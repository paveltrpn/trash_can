package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		deposit, yeild, approxErr float64
		duration                  int
	)

	fmt.Println("Enter deposit amount")
	fmt.Scan(&deposit)

	fmt.Println("Enter per month yeild percentage")
	fmt.Scan(&yeild)

	fmt.Println("Enter duration in years")
	fmt.Scan(&duration)

	for i := 0; i < duration*12; i++ {
		deposit += deposit * (yeild / 100)
		approxErr += (deposit*100 - math.Trunc(deposit*100)) / 100
		deposit = math.Trunc(deposit*100) / 100
	}

	fmt.Printf("Total yeild - %v\nWith approximation error - %v\n", deposit, approxErr)
}
