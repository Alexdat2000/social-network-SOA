package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var dest *string

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	targetURL := strings.TrimSuffix(*dest, "/") + r.URL.String()
	proxyReq, err := http.NewRequest(r.Method, targetURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}

	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	proxyResp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending proxy request: %v", err), http.StatusInternalServerError)
		return
	}
	defer proxyResp.Body.Close()

	respBodyBytes, err := io.ReadAll(proxyResp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading proxy response: %v", err), http.StatusInternalServerError)
		return
	}

	for key, values := range proxyResp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(proxyResp.StatusCode)
	w.Write(respBodyBytes)
}

func main() {
	listen := flag.String("listen", "localhost:8080", "listen for address")
	dest = flag.String("to", "localhost:8081", "redirect to this address")
	flag.Parse()

	log.Printf("Listening on %s", *listen)
	log.Printf("Redirecting to %s", *dest)
	http.HandleFunc("/", proxyHandler)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
