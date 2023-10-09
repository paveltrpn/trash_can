package main

// Задание 7. Игра "Угадай число"
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		value        int
		min          int = 1
		max          int = 10
		playerAnswer string
	)

	fmt.Println("Specify a integer value from 1 to 10")
	fmt.Scan(&value)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 4; i++ {
		guess := rand.Intn((max-min)+1) + min

		fmt.Printf("Maybe this number is - %v?\n", guess)
		fmt.Scan(&playerAnswer)

		if playerAnswer == "yes" {
			fmt.Println("Yo ho ho !!!")
			return
		}

		if playerAnswer == "more" {
			min = guess
			continue
		}

		if playerAnswer == "less" {
			max = guess
			continue
		}
	}
}
