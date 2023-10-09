// Задание 5. Решение квадратного уравнения
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		aVal, bVal, cVal, desc float64
	)

	fmt.Println("Enter A")
	fmt.Scan(&aVal)

	fmt.Println("Enter B")
	fmt.Scan(&bVal)

	fmt.Println("Enter C")
	fmt.Scan(&cVal)

	desc = math.Pow(bVal, 2.0) - 4*aVal*cVal

	if desc < 0 {
		fmt.Println("Equation does not have real roots!")
		return
	}

	if math.Abs(desc) < math.SmallestNonzeroFloat64 {
		root := (-bVal + math.Sqrt(desc)) / (2 * aVal)
		fmt.Printf("Equation has exactly one root - %v\n", root)
		return
	}

	if desc > 0 {
		root1 := (-bVal + math.Sqrt(desc)) / (2 * aVal)
		root2 := (-bVal - math.Sqrt(desc)) / (2 * aVal)
		fmt.Printf("Equation has two roots - %v and %v\n", root1, root2)
	}
}
