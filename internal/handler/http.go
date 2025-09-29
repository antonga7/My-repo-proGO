package handler

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"proGO/internal/fib"
	"strconv"
	"sync"
)

type Response struct {
	N   int      `json:"n"`
	Fib *big.Int `json:"fib"`
}

var fibCache = sync.Map{}

func WriteJSONError(w http.ResponseWriter, msg string, code int) {
	log.Printf("Ошибка [%d]: %s", code, msg)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос: %s %s", r.Method, r.URL.String())
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		WriteJSONError(w, "параметр 'n' обязателен", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		WriteJSONError(w, "параметр 'n' должен быть целым числом >= 0", http.StatusBadRequest)
		return
	}

	if val, ok := fibCache.Load(n); ok {
		fib := val.(*big.Int)
		resp := Response{N: n, Fib: fib}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Ошибка при кодировании ответа: %v", err)
		}

	}

	result, err := fib.Fibonacci(n)
	if err != nil {
		WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	fibCache.Store(n, result)
	resp := Response{N: n, Fib: result}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Ошибка при кодировании ответа: %v", err)
	}
	log.Printf("Успешно и сохранено в кэш: n=%d fib=%s", n, result.String())
}
