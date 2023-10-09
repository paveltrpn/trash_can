package main

import "fmt"

func main() {
	var (
		height int
	)

	fmt.Println("Enter pine height")
	fmt.Scan(&height)

	for i := 0; i < height*2; i++ {
		if (i % 2) == 1 {
			leafCnt := i % (height * 2)
			spaceCnt := (height - leafCnt/2) - 1

			for k := 0; k < spaceCnt; k++ {
				fmt.Print(" ")
			}

			for j := 0; j < leafCnt; j++ {
				fmt.Print("*")
			}
		}

		fmt.Printf("\n")
	}
}
