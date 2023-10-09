package main

import "fmt"

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func checkLuckySixDigits(ticket []int) bool {
	var (
		left, right int
	)

	for i := 0; i < 3; i++ {
		left += ticket[i]
		right += ticket[i+3]
	}

	if left == right {
		return true
	} else {
		return false
	}
}

func main() {
	var (
		amount int
	)

	for i := 100000; i < 999999; i++ {
		tmp := splitInt(i)

		if checkLuckySixDigits(tmp) {
			amount++
			// TEST
			//fmt.Printf("%v\n", i)
		}
	}

	fmt.Printf("Amount of lucky tickets is - %v\n", amount)
}
