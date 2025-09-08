package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

type Response struct {
	N   int      `json:"n"`
	Fib *big.Int `json:"fib"`
}

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

func writeJSONError(w http.ResponseWriter, msg string, code int) {
	log.Printf("Ошибка [%d]: %s", code, msg)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос: %s %s", r.Method, r.URL.String())
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		writeJSONError(w, "параметр 'n' обязателен", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		writeJSONError(w, "параметр 'n' должен быть целым числом >= 0", http.StatusBadRequest)
		return
	}

	result, err := fibonacci(n)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := Response{N: n, Fib: result}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Ошибка при кодировании ответа: %v", err)
	}

	log.Printf("Успешно: n=%d fib=%s", n, result.String())
}

func main() {
	http.HandleFunc("/fib", fibHandler)
	log.Println("Сервер запущен на хосте 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
