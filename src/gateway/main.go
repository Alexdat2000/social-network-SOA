package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"soa/gateway/api"
	"soa/gateway/content_grpc"
	"soa/gateway/stats_grpc"
)

func main() {
	server := api.Server{
		ContentAPI: content_grpc.InitContentClient("content:50051"),
		StatsAPI:   stats_grpc.InitStatsClient("stats:50052"),
	}
	r := chi.NewRouter()

	handler := api.HandlerFromMux(server, r)
	r.Mount("/", handler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
