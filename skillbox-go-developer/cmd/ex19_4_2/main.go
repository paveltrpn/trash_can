package main

import "fmt"

const (
	ARR_LENGTH = 6
)

var (
	arr = [ARR_LENGTH]int{5, 14, 0, 23, 3, 2}
)

func bubbleSort(arr *[ARR_LENGTH]int) {
	var (
		tmp int
	)

	for i := 0; i < ARR_LENGTH-1; {
		i++
		if arr[i] < arr[i-1] {
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp
			i = 0
		}
	}
}

func main() {
	bubbleSort(&arr)
	fmt.Println(arr)
}
