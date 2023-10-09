package main

import "fmt"

func main() {
	var (
		price, discount float64
	)

	fmt.Printf("Enter price\n")
	fmt.Scan(&price)

	fmt.Printf("Enter discount (in percents)\n")
	fmt.Scan(&discount)

	if discount > 30 {
		fmt.Printf("Discount is too big!\n")
		return
	}

	discountValue := price * (discount / 100)

	if discountValue > 2000 {
		fmt.Printf("Discout amount is too big!\n")
		return
	} else {
		fmt.Printf("Discount value is - %v%%\nDiscount amount is - %v\n", discount, discountValue)
	}
}
