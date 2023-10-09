package main

import (
	"fmt"
	"strconv"
	"strings"
)

type simpleQueue struct {
	queue []int
	index int
}

func buildSimpleQueue(q []int) simpleQueue {
	var rt simpleQueue

	rt.queue = make([]int, len(q))
	copy(rt.queue, q)
	rt.index = 0

	return rt
}

func (q *simpleQueue) next() (bool, int) {
	if q.index == len(q.queue) {
		return false, 0
	}
	var rt int = q.queue[q.index]
	q.index++
	return true, rt
}

func lemonadeChange(bills []int) bool {
	var (
		five, ten, twenty int
	)

	billsQueue := buildSimpleQueue(bills)

	for {
		serve, bill := billsQueue.next()

		if !serve {
			fmt.Printf("Queue is empty!\n")
			return true
		}

		switch bill {
		case 5:
			five++
			fmt.Printf("You give 5, then take your icecream!!!\n")
			continue
		case 10:
			if five < 1 {
				fmt.Printf("You give 10, but unfurtunatly i don't have change(((\n")
				return false
			} else {
				fmt.Printf("You give 10, i give five and your icecream!!!\n")
				ten++
				five--
				continue
			}
		case 20:
			if five >= 3 {
				fmt.Printf("You give 20, i give you three fives and your icecream!!!\n")
				twenty++
				five -= 3
				continue
			}

			if (ten >= 1) && (five >= 1) {
				fmt.Printf("You give 20, i give you ten, five and your icecream!!!\n")
				twenty++
				ten--
				five--
				continue
			}
			fmt.Printf("You give 20, but unfurtunatly i don't have change(((\n")
			return false
		}
	}
}

func main() {
	var (
		cashString string
	)

	fmt.Println("Enter byers cash queue in format - \"[5,5,10,20]\", etc.")
	fmt.Scan(&cashString)

	bills := strings.Split(strings.Trim(cashString, "[]"), ",")

	var billsInt []int = make([]int, len(bills))

	for i, b := range bills {
		if (b != "5") && (b != "10") && (b != "20") {
			fmt.Printf("Wrong bill - %v\n", b)
			return
		}

		billsInt[i], _ = strconv.Atoi(b)
	}

	fmt.Printf("Is queue served - %v\n", lemonadeChange(billsInt))
}
