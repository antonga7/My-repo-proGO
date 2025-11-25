package handler

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"proGO/internal/fib"
	"proGO/internal/storage"
	"strconv"
)

type Handler struct {
	store storage.Storer
}

func NewHandler(s storage.Storer) *Handler {
	return &Handler{store: s}
}

type Response struct {
	N   int      `json:"n"`
	Fib *big.Int `json:"fib"`
}

func WriteJSONError(w http.ResponseWriter, msg string, code int) {
	log.Printf("Ошибка [%d]: %s", code, msg)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
		http.Error(w, "Не удалось закодировать ответ", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FibHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx := context.Background()

	if val, err := h.store.Get(ctx, n); err != nil {
		WriteJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	} else if val != nil {
		resp := Response{N: n, Fib: val}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w,"Не удалось закодировать ответ", http.StatusInternalServerError)
			return
		}
		log.Printf("Извлечено из базы: n=%d, fib =%s", n, val.String())
		return
	}

	result, err := fib.Fibonacci(n)
	if err != nil {
		WriteJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.store.Set(ctx, n, result); err != nil {
		log.Printf("Не удалось сохранить в базу: %v", err)
	}

	resp := Response{N: n, Fib: result}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Ошибка при кодировании ответа: %v", err)
	}
	log.Printf("Успешно и сохранено в базу: n=%d fib=%s", n, result.String())
}
