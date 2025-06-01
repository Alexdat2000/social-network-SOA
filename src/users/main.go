package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"soa/users/api"
)

func main() {
	server := api.Server{
		DB:       api.InitDB(),
		Kafka:    api.InitKafka(),
		Handlers: api.InitAuthHandlers(),
	}
	r := chi.NewRouter()

	handler := api.HandlerFromMux(server, r)
	r.Mount("/", handler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
