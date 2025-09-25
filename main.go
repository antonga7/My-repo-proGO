package main

import (
	"log"
	"net/http"
	"proGO/internal/handler"
)

func main() {
	http.HandleFunc("/fib", handler.FibHandler)
	log.Println("Сервер запущен на хосте 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
