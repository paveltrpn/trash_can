package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	ARR_SIZE = 10
)

var (
	arr []int
)

func binSearchMostLeft(arr []int, elem int) int {
	var (
		index = -1
		min   = 0
		max   = len(arr) - 1
	)

	for max >= min {
		middle := (max + min) / 2
		if arr[middle] == elem {
			index = middle
			for (middle > 0) && (arr[middle] == elem) {
				middle--

				if arr[middle] != elem {
					middle++
					return middle
				}
			}
			break
		} else if arr[middle] > elem {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	return index
}

func main() {
	var (
		input int
	)

	arr = make([]int, ARR_SIZE)

	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < ARR_SIZE; i++ {
		arr[i] = rand.Intn(10 * ARR_SIZE)
	}

	sort.Ints(arr)

	fmt.Println(arr)

	fmt.Println("Enter number")
	fmt.Scan(&input)

	fmt.Println(binSearchMostLeft(arr, input))
}
