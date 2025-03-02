/*
 * Social Network
 *
 * This is a homework project for Service-Oriented Architectures cource
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	sw "soa/users/go"
)

func main() {
	sw.InitDB()
	sw.InitAuthHandler()
	log.Printf("Server started")
	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
