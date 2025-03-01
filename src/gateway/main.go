package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var sourcePort = 8080
var destPort = 8081

func handler(w http.ResponseWriter, r *http.Request) {
	targetURL := fmt.Sprintf("localhost:%d", destPort)

	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Failed to create proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to read answer", http.StatusInternalServerError)
	}
}

func main() {
	sourcePort = *flag.Int("from", 8080, "from port")
	destPort = *flag.Int("to", 8081, "to port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", sourcePort), nil))
}
