package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"time"
)

var dest *string

func handler(w http.ResponseWriter, r *http.Request) {
	targetURL := *dest

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
	listen := flag.String("listen", "localhost:8080", "listen for address")
	dest = flag.String("to", "localhost:8081", "redirect to this address")
	flag.Parse()

	log.Printf("Listening on %s", *listen)
	log.Printf("Redirecting to %s", *dest)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
