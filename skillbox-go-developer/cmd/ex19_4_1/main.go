package main

import (
	"fmt"
	"sort"
)

var (
	arrOne = [...]int{5, 400, 12, 45}
	arrTwo = [...]int{9, 105, 3, 98, 901}
	merged [9]int
)

func main() {
	var (
		i, j, end int
	)

	sort.Ints(arrOne[:])
	sort.Ints(arrTwo[:])

	i = len(arrOne) - 1
	j = len(arrTwo) - 1
	end = len(arrOne) + len(arrTwo) - 1

	for j >= 0 {
		if i >= 0 && arrOne[i] > arrTwo[j] {
			merged[end] = arrOne[i]
			i--
		} else {
			merged[end] = arrTwo[j]
			j--
		}
		end--
	}

	fmt.Println(merged)
}
