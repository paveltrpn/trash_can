// Задание 5. Ресторан
package main

import "fmt"

func main() {
	var (
		day, guestCount, value int
		discount, incrase      int
	)

	fmt.Println("Введите день недели")
	fmt.Scan(&day)

	fmt.Println("Введите число гостей")
	fmt.Scan(&guestCount)

	fmt.Println("Введите сумму чека")
	fmt.Scan(&value)

	if day == 1 {
		discount = value / 10
		fmt.Println("Скидка по понедельникам:", discount)
	}

	if (day == 5) && (value > 10000) {
		discount = (value * 5) / 100
		fmt.Println("Скидка по пятницам:", discount)
	}

	if guestCount > 7 {
		incrase = value / 10
		fmt.Println("Надбавка за обслуживание:", incrase)
	}

	price := value + discount - incrase
	fmt.Println("Сумма к оплате:", price)
}
