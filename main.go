// Запускает сервер с поддержкой базы данных
package main

import (
	"log"
	"net/http"
	"proGO/internal/handler"
	"proGO/internal/storage/postgres"
)

func main() {
	connStr := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"

	store, err := postgres.New(connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе: %v", err)
	}

	h := handler.NewHandler(store)

	http.HandleFunc("/fib", h.FibHandler)
	log.Println("Сервер запущен на хосте 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
