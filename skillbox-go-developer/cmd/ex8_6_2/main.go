package main

import "fmt"

func main() {
	var (
		day       string
		weekWork  = [...]string{"понедельник", "вторник", "среда", "четверг", "пятница"}
		dayNumber int
	)

	fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
	fmt.Scan(&day)

	switch day {
	case "пн":
		dayNumber = 0
	case "вт":
		dayNumber = 1
	case "ср":
		dayNumber = 2
	case "чт":
		dayNumber = 3
	case "пт":
		dayNumber = 4
	default:
		fmt.Println("Кокой-то странный день недели")
		return
	}

	for i := dayNumber; i < 5; i++ {
		fmt.Println(weekWork[i])
	}
}
