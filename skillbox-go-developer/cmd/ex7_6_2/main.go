package main

import "fmt"

func main() {
	var (
		rows, columns int
		checker       = map[int]string{0: " ", 1: "*"}
	)

	fmt.Println("Enter columns count")
	fmt.Scan(&columns)

	fmt.Println("Enter rows count")
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%v", checker[(j+i)%2])
		}
		fmt.Printf("\n")
	}
}
