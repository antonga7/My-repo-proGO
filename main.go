// main package starts the proGO-Fibonacci server,
// processes routes via chi and oapi-codegen.
package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"proGO/api"
	"proGO/internal/handler"
)

func main() {
	r := chi.NewRouter()
	fibHandler := &handler.FibHandler{}
	api.HandlerFromMux(
		api.NewStrictHandler(fibHandler, nil),
		r,
	)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
