// Задание 1. Определение координатной плоскости точки
package main

import "fmt"

func main() {
	var (
		xCrd, yCrd int
	)

	fmt.Println("Введите значение координаты Х")
	fmt.Scan(&xCrd)

	fmt.Println("Введите значение координаты Y")
	fmt.Scan(&yCrd)

	if (xCrd == 0) || (yCrd == 0) {
		fmt.Println("Некорректный ввод!")
		return
	}

	if (xCrd > 0) && (yCrd > 0) {
		fmt.Println("Точка находится в первой четверти")
		return
	}

	if (xCrd < 0) && (yCrd > 0) {
		fmt.Println("Точка находится во второй четверти")
		return
	}

	if (xCrd < 0) && (yCrd < 0) {
		fmt.Println("Точка находится в третьей четверти")
		return
	}

	if (xCrd > 0) && (yCrd < 0) {
		fmt.Println("точка находится в четвёртой четверти")
		return
	}
}
