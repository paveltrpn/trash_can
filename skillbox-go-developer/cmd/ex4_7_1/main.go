// Задание 1. Баллы ЕГЭ
package main

import "fmt"

func main() {
	var (
		firstEx, secondEx, thirdEx int
	)

	fmt.Println("Баллы ЕГЭ.")

	fmt.Println("Введите результат первого экзамена")
	fmt.Scan(&firstEx)

	fmt.Println("Введите результаты второго экзамена")
	fmt.Scan(&secondEx)

	fmt.Println("Введите результаты третьего экзамена")
	fmt.Scan(&thirdEx)

	if (firstEx < 0) || (secondEx < 0) || (thirdEx < 0) {
		fmt.Println("Одно из введённых значений отрицательно!")
		return
	}

	sum := firstEx + secondEx + thirdEx

	fmt.Println("Сумма проходных баллов: 275")
	fmt.Printf("Количество набранных баллов: %v\n", sum)

	if sum < 275 {
		fmt.Println("Количество набранных баллов: 275\nYou should not pass!")
	} else {
		fmt.Println("Вы поступили!!!")
	}
}
