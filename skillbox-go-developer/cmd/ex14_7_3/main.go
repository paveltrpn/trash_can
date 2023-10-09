package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const (
	SUM = 10
	MLT = 20
)

type Number interface {
	constraints.Integer | constraints.Float
}

func sum[T Number](num T) (rt T) {
	rt = num + T(SUM)
	return
}

func mlt[T Number](num T) (rt T) {
	rt = num * T(MLT)
	return
}

func main() {
	var (
		foo float64
	)

	fmt.Println("Enter a number!")
	fmt.Scan(&foo)

	foo = sum(foo)
	foo = mlt(foo)

	fmt.Printf("mapped number = %v\n", foo)
}
