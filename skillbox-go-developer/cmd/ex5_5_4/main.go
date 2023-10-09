package main

import "fmt"

func main() {
	var (
		price, firstCoin, secondCoin, thirdCoin int
	)

	fmt.Println("Enter price")
	fmt.Scan(&price)

	if price <= 0 {
		fmt.Println("Incorrect price!")
		return
	}

	fmt.Println("Enter first coin nominal")
	fmt.Scan(&firstCoin)

	fmt.Println("Enter second coin nominal")
	fmt.Scan(&secondCoin)

	fmt.Println("Enter third coin nominal")
	fmt.Scan(&thirdCoin)

	if (price == firstCoin) || (price == secondCoin) || (price == thirdCoin) {
		fmt.Println("Price is equal to one of coins!")
		return
	}

	if price == (firstCoin + secondCoin) {
		fmt.Println("Price is equal to sum of first and second coins!")
		return
	}

	if price == (firstCoin + thirdCoin) {
		fmt.Println("Price is equal to sum of first and third coins!")
		return
	}

	if price == (secondCoin + thirdCoin) {
		fmt.Println("Price is equal to sum of second and third coins!")
		return
	}

	if price == (firstCoin + secondCoin + thirdCoin) {
		fmt.Println("Prise is equal to sum of all three coins!")
		return
	}

	fmt.Println("You cannot pay by any of coins or it's combination!")
}
