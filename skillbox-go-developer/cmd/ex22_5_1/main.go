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

func find(arr [ARR_SIZE]int, elem int) int {
	var index = -1
	for i := 0; i < ARR_SIZE; i++ {
		if arr[i] == elem {
			index = i
			break
		}
	}

	return index
}

func main() {
	var (
		input int
	)

	rand.Seed(time.Now().UnixMicro())
	for i := 0; i < ARR_SIZE; i++ {
		arr[i] = rand.Intn(10 * ARR_SIZE)
	}

	fmt.Println(arr)

	fmt.Println("Enter number")
	fmt.Scan(&input)

	id := find(arr, input)
	if id != -1 {
		fmt.Printf("Element - %v, placed at index - %v with elapsed elements - %v\n", input, id, (ARR_SIZE-id)-1)
	} else {
		fmt.Println(0)
	}
}
