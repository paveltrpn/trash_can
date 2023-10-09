// Задание 4. Три числа: ещё попытка
package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int

	arr = make([]int, 3)

	fmt.Println("Введите первое число")
	fmt.Scan(&arr[0])

	fmt.Println("Введите второе число")
	fmt.Scan(&arr[1])

	fmt.Println("Введите третье число")
	fmt.Scan(&arr[2])

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for i := 0; i < 3; i++ {
		if arr[i] >= 5 {
			fmt.Printf("Среди введённых чисел %v больше или равны 5\n", 3-i)
			return
		}
	}

	fmt.Println("Среди введённых чисел нет чисел больше или равных 5")
}
