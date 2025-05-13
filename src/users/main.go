package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"soa/users/api"
)

func main() {
	api.InitDB()
	api.InitAuthHandler()
	api.ConnectToKafka()

	server := api.Server{}
	r := chi.NewRouter()

	handler := api.HandlerFromMux(server, r)
	r.Mount("/", handler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
