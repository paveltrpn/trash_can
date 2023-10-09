package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ARR_SIZE = 15
)

func main() {

	arr := make([]int, ARR_SIZE)

	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < ARR_SIZE; i++ {
		arr[i] = rand.Intn(10 * ARR_SIZE)
	}

	bubble := func(arr []int, cmp func(int, int) bool) {
		var (
			tmp int
		)

		for i := 0; i < len(arr)-1; {
			i++
			if cmp(arr[i], arr[i-1]) {
				tmp = arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = tmp
				i = 0
			}
		}

	}

	reverse := func(arr []int) {
		for i, j := 0, len(arr)-1; i != j; {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
			i++
			j--
		}
	}

	fmt.Println(arr)

	bubble(arr, func(a, b int) bool {
		if a > b {
			return true
		} else {
			return false
		}
	})

	reverse(arr)
	reverse(arr)

	fmt.Println(arr)
}
