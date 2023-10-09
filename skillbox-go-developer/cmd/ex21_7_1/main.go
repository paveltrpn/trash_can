package main

import (
	"fmt"
	"math"
)

func main() {
	foo := func(x int16, y uint8, z float32) float32 {
		return 2*float32(x) + float32(math.Pow(float64(y), 2)) - 3/z
	}(10, 30, 14.1)
	fmt.Println(foo)
}
