// Задание 2. Три числа
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

	if arr[len(arr)-1] > 5 {
		fmt.Println("Среди введённых чисел есть число больше 5")
	} else {
		fmt.Println("Среди введённых чисел нет числа больше 5")
	}
}
