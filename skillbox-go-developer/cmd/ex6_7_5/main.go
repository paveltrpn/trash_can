package main

import "fmt"

func main() {
	var (
		firstBskt, secondBskt, thirdBskt int
		firstCap, secondCap, thirdCap    int
		isAppleInHand                    bool
	)

	fmt.Printf("Enter fisrst capacity\n")
	fmt.Scan(&firstCap)

	fmt.Printf("Enter second capacity\n")
	fmt.Scan(&secondCap)

	fmt.Printf("Enter third capacity\n")
	fmt.Scan(&thirdCap)

	for {
		isAppleInHand = true

		if (isAppleInHand) && (firstBskt < firstCap) {
			firstBskt++
			isAppleInHand = false
			continue
		}

		if (isAppleInHand) && (secondBskt < secondCap) {
			secondBskt++
			isAppleInHand = false
			continue
		}

		if (isAppleInHand) && (thirdBskt < thirdCap) {
			thirdBskt++
			isAppleInHand = false
			continue
		}

		if (firstBskt == firstCap) && (secondBskt == secondCap) && (thirdBskt == thirdCap) {
			break
		}
	}

	fmt.Printf("firstBskt - %v\nsecondBsct - %v\nthirdBsct - %v\n", firstBskt, secondBskt, thirdBskt)
}
