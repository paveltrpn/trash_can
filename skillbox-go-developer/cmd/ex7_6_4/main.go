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
		currentDistance, greatestDistance int
		nowTicket, prevTicket             int
	)

	prevTicket = 100000

	for i := 100000; i < 999999; i++ {
		tmp := splitInt(i)

		nowTicket = i

		if checkLuckySixDigits(tmp) {

			currentDistance = nowTicket - prevTicket
			// TEST
			// fmt.Printf("dst - %v, now - %v, prev - %v\n", currentDistance, nowTicket, prevTicket)

			if (nowTicket != prevTicket) && (currentDistance > greatestDistance) {
				greatestDistance = currentDistance
				// TEST
				//fmt.Printf("gst - %v, now - %v, prev - %v\n", greatestDistance, nowTicket, prevTicket)
				prevTicket = nowTicket
			}
			//prevTicket = i
		}
	}

	fmt.Printf("Greatest distance between lucky tickets is - %v\n", greatestDistance)
}
