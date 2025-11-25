package handler

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// FakeStore реализует интерфейс Storer для целей тестирования
type FakeStore struct {
	data map[int]*big.Int
}

// Get - фиктивная реализация метода Get
func (f *FakeStore) Get(_ context.Context, n int) (*big.Int, error) {
	if val, ok := f.data[n]; ok {
		return val, nil
	}
	return nil, nil
}

// Get - фиктивная реализация метода Set
func (f *FakeStore) Set(_ context.Context, n int, val *big.Int) error {
	f.data[n] = val
	return nil
}

func TestFibHandler(t *testing.T) {
	store := &FakeStore{data: make(map[int]*big.Int)}
	h := NewHandler(store)
	req := httptest.NewRequest(http.MethodGet, "/fibn?n=10", nil)
	rec := httptest.NewRecorder()

	h.FibHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("ожидался статус %d, а получен %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"fib":55`) {
		t.Errorf("ошибка, ожидалось другое значение: %s", body)
	}

	if _, ok := store.data[10]; !ok {
		t.Errorf("значение не было сохранено в БД")
	}
}
