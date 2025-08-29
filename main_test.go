package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	type TestCase struct {
		input    int
		expected int
	}
	testCases := []TestCase{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 2, expected: 1},
		{input: 3, expected: 2},
		{input: 4, expected: 3},
		{input: 5, expected: 5},
		{input: 6, expected: 8},
		{input: 10, expected: 55},
		{input: -5, expected: 5},
		{input: 93, expected: 1220016041512187673},
		{input: 92, expected: 7540113804746346429},
	}

	for _, tc := range testCases {
		// Вызываем нашу функцию с входными данными из тестового случая
		result := fibonacci(tc.input)

		// ТВОЯ ЗАДАЧА: сравни result и tc.expected
		if result != tc.expected {
			t.Errorf("fibonacci(%d) вернул %d, а ожидалось %d", tc.input, result, tc.expected)
		}
	}
}
