package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		foo                                uint8  = math.MaxUint8
		bar                                uint16 = math.MaxUint16
		fooOverloadCount, barOverloadCount int
	)

	var i uint32
	for i = 0; i < math.MaxUint32; i++ {
		foo++
		bar++

		if foo == 0 {
			fooOverloadCount++
		}

		if bar == 0 {
			barOverloadCount++
		}
	}

	fmt.Printf("uint8 overloads count - %v\nuint16 overloads count - %v\n", fooOverloadCount, barOverloadCount)
}
