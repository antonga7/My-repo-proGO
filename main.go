package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func fibonacci(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("нельзя вычислить число Фибоначчи для отрицательного числа")
	}

	if n == 0 {
		return big.NewInt(0), nil
	}

	if n == 1 {
		return big.NewInt(1), nil
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		sum := new(big.Int).Add(a, b)
		a, b = b, sum
	}
	return b, nil
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
		result, err := fibonacci(num)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Число Фибоначчи, соотвествующее введенному номеру: ", result)
		break
	}
}
