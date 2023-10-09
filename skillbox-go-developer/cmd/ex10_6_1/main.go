package main

import (
	"fmt"
	"math"
)

func fctrl(n int) float64 {
	var rt float64 = 1

	for i := 1; i <= n; i++ {
		rt *= float64(i)
	}

	return rt
}

func main() {
	var (
		x, eps        float64
		n             int
		euler, peuler float64 = 0.0, 0.0
	)

	fmt.Println("Enter exponent - x")
	fmt.Scan(&x)

	fmt.Println("Enter accurancy factor - n")
	fmt.Scan(&n)

	eps = math.Pow(0.1, float64(n))

	for i := 0; ; i++ {
		euler += math.Pow(x, float64(i)) / fctrl(i)
		if math.Abs(euler-peuler) < eps {
			break
		}
		peuler = euler
	}

	fmt.Println(euler)
}
