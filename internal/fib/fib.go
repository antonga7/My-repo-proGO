package fib

import (
	"errors"
	"math/big"
)

func Fibonacci(n int) (*big.Int, error) {
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
