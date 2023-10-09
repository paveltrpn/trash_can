package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ARR_SIZE = 10
)

var (
	arr [ARR_SIZE]int
)

func insertionsort(items *[ARR_SIZE]int) {
	var n = ARR_SIZE
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < ARR_SIZE; i++ {
		arr[i] = rand.Intn(10 * ARR_SIZE)
	}

	fmt.Println(arr)
	insertionsort(&arr)
	fmt.Println(arr)
}
