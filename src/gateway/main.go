package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"soa/gateway/api"
	pb "soa/gateway/content_grpc"
)

func main() {
	server := api.Server{
		ContentAPI: pb.InitContentClient("localhost:8082"),
	}
	r := chi.NewRouter()

	handler := api.HandlerFromMux(server, r)
	r.Mount("/", handler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
