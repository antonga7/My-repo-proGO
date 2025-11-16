package handler

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type FakeStore struct {
	data map[int]*big.Int
}

func (f *FakeStore) Get(ctx context.Context, n int) (*big.Int, error) {
	if val, ok := f.data[n]; ok {
		return val, nil
	}
	return nil, nil
}

func (f *FakeStore) Set(ctx context.Context, n int, val *big.Int) error {
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
