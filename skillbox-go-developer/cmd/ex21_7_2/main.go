package main

import "fmt"

type fnc func(int, int) int

func caller(f fnc) {
	defer f(1001, 9)
}

func foo(x, y int) int {

	fmt.Println("i'am foo")
	return x + y
}

func bar(x, y int) int {
	fmt.Println("I'am bar")
	return x - y
}

func baz(x, y int) int {
	fmt.Println("I'am baz, and i return 0")
	return 0
}

func main() {
	caller(foo)
	caller(bar)
	caller(baz)
}
