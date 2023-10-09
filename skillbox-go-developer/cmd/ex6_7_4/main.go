package main

import "fmt"

func main() {
	var (
		ten     int
		hundred int
		tousand int
	)

	for i := 0; i < 1000; i++ {
		if ten < 10 {
			ten++
		}

		if hundred < 100 {
			hundred++
		}

		if tousand < 1000 {
			tousand++
		}
	}

	fmt.Printf("%v %v %v", ten, hundred, tousand)
}
