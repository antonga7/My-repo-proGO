package integration_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"net"
	"net/http"
	"proGO/api"
	"proGO/internal/handler"
	dbpostgres "proGO/internal/storage/postgres"
	"sync"
	"testing"
	"time"
)

func TestFibAPI(t *testing.T) {
	ctx := context.Background()

	t.Log("Тест запустился")

	// 1. Поднимаем контейнер PostgresSQL
	pgContainer, err := tcpostgres.Run(
		ctx,
		"postgres:16-alpine",
		tcpostgres.WithDatabase("testdb"),
		tcpostgres.WithUsername("testuser"),
		tcpostgres.WithPassword("testpass"),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer pgContainer.Terminate(ctx)

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("postgres запущен и доступен по:", connStr)

	// 2. Подключаемся к базе через Store
	store, err := dbpostgres.New(connStr)
	if err != nil {
		t.Fatal(err)
	}

	// 3. Ждем готовность Postgres выполнять запросы(у меня здесь тест крашился ранее)
	maxAttempts := 10
	for i := 0; i < maxAttempts; i++ {
		_, err := store.Pool().Exec(ctx, "SELECT 1")
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
		if i == maxAttempts-1 {
			t.Fatal("Postgres не готов к соединению:", err)
		}
	}

	// 4. Создаем таблицу fibonacci
	_, err = store.Pool().Exec(ctx, `CREATE TABLE IF NOT EXISTS fibonacci (n BIGINT PRIMARY KEY, value TEXT)`)
	if err != nil {
		t.Fatal("Не удалось создать таблицу", err)
	}

	// 5. Создаём HTTP-сервер
	fibHandler := handler.NewFibHandler(store)
	r := chi.NewRouter()
	api.HandlerFromMux(api.NewStrictHandler(fibHandler, nil), r)

	addr := "localhost:8081"
	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	var mu sync.Mutex
	mu.Lock()

	go func() {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}

		mu.Unlock()

		if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Fatal(err)
		}
	}()

	defer server.Close()

	// 6. HTTP запрос
	resp, err := http.Get("http://" + addr + "/fib?n=10")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Ожидался 200 получился %d", resp.StatusCode)
	}

	// 7. JSON и проверка результата
	var result struct {
		Result string `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if result.Result != "55" {
		t.Fatalf("ожидалось 55 получилось %s", result.Result)
	}
}
