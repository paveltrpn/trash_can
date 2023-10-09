package main

import "fmt"

func main() {
	var (
		month string
	)

	fmt.Println("Введите месяц года")
	fmt.Scan(&month)

	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Время года - зима")
	case "март", "апрель", "май":
		fmt.Println("Время года - весна")
	case "июнь", "июль", "август":
		fmt.Println("Время года - лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Время года - осень")
	default:
		fmt.Println("Такого месяца не существует!")
	}
}
