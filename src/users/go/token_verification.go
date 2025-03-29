package api

import (
	"fmt"
	"log"
	"net/http"
)

func TokenGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jwt := r.URL.Query().Get("jwt")
	if jwt == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name, err := ValidateToken(jwt)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var totalCount int
	err = DB.QueryRow("SELECT COUNT(*) FROM users where username=$1", name).Scan(&totalCount)
	if totalCount == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\"%s\"", name)
}
