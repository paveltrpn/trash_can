package main

import (
	"fmt"
	"log"
	"strconv"
)

func square(ch chan int) chan int {
	rt := make(chan int)

	nbr := <-ch
	fmt.Printf("square = %v\n", nbr*nbr)

	go func() {
		rt <- nbr * nbr
	}()

	return rt
}

func mult(ch chan int) {
	nbr := <-ch
	fmt.Printf("dubbled = %v\n", nbr*2)
}

func main() {
	var (
		inp string
	)

	nbrCh := make(chan int)
	sqCh := make(chan int)

	for {
		fmt.Scan(&inp)
		if inp == "stop" {
			fmt.Println("stop word recieved!")
			break
		}

		number, err := strconv.Atoi(inp)
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			nbrCh <- number
		}()

		sqCh = square(nbrCh)

		mult(sqCh)
	}
}
