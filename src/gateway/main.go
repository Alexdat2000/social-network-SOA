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

	"soa/gateway/content"
)

func proxyHandler(w http.ResponseWriter, r *http.Request, dest string) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	targetURL := strings.TrimSuffix(dest, "/") + r.URL.String()
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
	users := flag.String("users", "localhost:8081", "user server at this address")
	cont := flag.String("content", "localhost:8082", "content server at this address")
	flag.Parse()

	content.InitGrpc(*cont)

	log.Printf("Listening on %s", *listen)
	log.Printf("User server at %s", *users)
	log.Printf("Content server at %s", *cont)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) { proxyHandler(w, r, *users) })
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { proxyHandler(w, r, *users) })
	http.HandleFunc("/entry", func(w http.ResponseWriter, r *http.Request) { content.HandleEntry(w, r, *users, *cont) })
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) { content.HandleList(w, r, *users, *cont) })
	log.Fatal(http.ListenAndServe(*listen, nil))
}
