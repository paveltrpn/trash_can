// Задание 3. Проверить есть ли совпадающие числа
package main

import "fmt"

func main() {
	var (
		first, second, third int
	)

	fmt.Println("Enter first")
	fmt.Scan(&first)

	fmt.Println("Enter second")
	fmt.Scan(&second)

	fmt.Println("Enter third")
	fmt.Scan(&third)

	if (first == second) || (second == third) || (first == third) {
		fmt.Println("At least two of entered numbers equal!")
		return
	}

	fmt.Println("Therse no equal numbers!")
}
