package main

import "fmt"

var (
	arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func reverse(a []int) []int {
	rt := make([]int, len(a))

	for i, j := len(a)-1, 0; i >= 0; i, j = i-1, j+1 {
		rt[i] = a[j]
	}

	return rt
}

func main() {
	fmt.Println(reverse(arr[:]))
}
