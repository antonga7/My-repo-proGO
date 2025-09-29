package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFibHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/fib?n=10", nil)
	rec := httptest.NewRecorder()
	FibHandler(rec, request)
	if rec.Code != http.StatusOK {
		t.Fatalf("ожидался статус %d, а получен %d", http.StatusOK, rec.Code)
	}

	// 5. Проверяем тело ответа
	body := rec.Body.String()
	if !strings.Contains(body, `"fib":55`) {
		t.Errorf("ошибка, ожидалось другое значение: %s", body)
	}

}
