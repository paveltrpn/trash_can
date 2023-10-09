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
	arr []int
)

func isOdd(num int) bool {
	if (num % 2) == 1 {
		return true
	} else {
		return false
	}
}

func oddEvenSieve(a []int) (odd []int, even []int) {
	for _, n := range a {
		if isOdd(n) {
			odd = append(odd, n)
		} else {
			even = append(even, n)
		}
	}
	return
}

func main() {

	arr = make([]int, ARR_SIZE)

	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < ARR_SIZE; i++ {
		arr[i] = rand.Intn(10 * ARR_SIZE)
	}

	var (
		odd, even []int
	)

	odd, even = oddEvenSieve(arr)

	fmt.Println(arr)
	fmt.Println(odd)
	fmt.Println(even)
}
