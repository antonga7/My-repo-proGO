package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fibonacci(n int) int {
	if n <= 1 && n >= 0 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	var input string

	for {
		fmt.Print("Введите целое число (0 или больше): ")
		fmt.Scanln(&input)

		// Убираем пробелы
		input = strings.TrimSpace(input)

		// Пробуем преобразовать в int
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число без дробной части.")
			continue
		}

		if num < 0 {
			fmt.Println("Ошибка: число должно быть 0 или больше.")
			continue
		}

		fmt.Printf("Число Фибоначчи, соотвествующее введенному номеру: %d.\n", fibonacci(num))
		break
	}
}
