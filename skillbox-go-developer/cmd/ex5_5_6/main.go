// Задача 6. Счастливый билет
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

func main() {
	var (
		ticket int
	)

	fmt.Println("Enter ticket number")
	fmt.Scan(&ticket)

	if (ticket < 1000) && (ticket > 9999) {
		fmt.Println("Incorrect ticket number!")
		return
	}

	nmbrs := splitInt(ticket)

	if (nmbrs[0] == nmbrs[3]) && (nmbrs[1] == nmbrs[2]) {
		fmt.Println("It's a reflect number!")
		return
	}

	if (nmbrs[0] + nmbrs[1]) == (nmbrs[2] + nmbrs[3]) {
		fmt.Println("It's lucky ticket!")
	} else {
		fmt.Println("It's regular ticket!")
	}
}
