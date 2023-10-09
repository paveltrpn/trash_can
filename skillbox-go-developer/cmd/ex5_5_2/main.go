// Задание 2. Проверить, сто одно из чисел - положительное
package main

import "fmt"

func main() {
	var (
		first, second, third int
	)

	fmt.Println("Введите первое число")
	fmt.Scan(&first)

	fmt.Println("Введите второе число")
	fmt.Scan(&second)

	fmt.Println("Введите третье число")
	fmt.Scan(&third)

	if (first > 0) || (second > 0) || (third > 0) {
		fmt.Println("Хотя бы одно из чисел больше ноля")
	} else {
		fmt.Println("Среди введённых чисел нет положительных")
	}

}
