package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	type TestCase struct {
		input    int
		expected string
	}
	testCases := []TestCase{
		{input: 0, expected: "0"},
		{input: 1, expected: "1"},
		{input: 2, expected: "1"},
		{input: 3, expected: "2"},
		{input: 4, expected: "3"},
		{input: 5, expected: "5"},
		{input: 6, expected: "8"},
		{input: 10, expected: "55"},
		{input: -5, expected: ""},
		{input: 93, expected: "12200160415121876738"},
		{input: 92, expected: "7540113804746346429"},
		{input: 100, expected: "354224848179261915075"},
	}

	for _, tc := range testCases {
		// Вызываем нашу функцию с входными данными из тестового случая
		result, err := fibonacci(tc.input)
		if err != nil {
			if tc.input < 0 {
				continue
			}
			t.Errorf("неожиданная ошибка: %v", err)
		}

		// ТВОЯ ЗАДАЧА: сравни result и tc.expected
		if result.String() != tc.expected {
			t.Errorf("fibonacci(%d) вернул %s, а ожидалось %s", tc.input, result.String(), tc.expected)
		}
	}
}
