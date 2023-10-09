package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		foo, bar int16
	)

	fmt.Println("Enter first number")
	fmt.Scan(&foo)

	fmt.Println("Enter second number")
	fmt.Scan(&bar)

	var mult int32 = int32(foo) * int32(bar)

	fmt.Println(mult)

	switch {
	case (mult > 0) && (mult < 255):
		fmt.Println("uint8")
	case (mult > -127) && (mult < 0):
		fmt.Println("int8")

	case (mult > 255) && (mult < math.MaxUint16):
		fmt.Println("uint16")
	case (mult > math.MinInt16) && (mult < 0):
		fmt.Println("int16")

	case (mult > math.MaxUint16):
		fmt.Println("uint32")
	case (mult > math.MinInt32) && (mult < 0):
		fmt.Println("int32")
	}
}
