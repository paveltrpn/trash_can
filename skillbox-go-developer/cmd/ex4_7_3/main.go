// Задание 3. Банкомат
package main

import (
	"fmt"
)

func main() {
	var value int

	fmt.Println("Введите первое число")
	fmt.Scan(&value)

	if (value > 100000) || (value <= 0) {
		fmt.Println("Данную сумму невозможно выдать! В банкомате недостаточно денег или введена некорректная сумма!")
		return
	}

	tmp := value % 100
	if tmp != 0 {
		fmt.Println("Данную сумму невозможно выдать! Банкомат выдаёт только купюры по 100 рублей!")
		return
	}

	fmt.Printf("Вы сняли %v рублей\n", value)
}
