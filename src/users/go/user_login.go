package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func LoginGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	password := r.URL.Query().Get("password")
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(username) > 32 || len(password) > 32 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := fmt.Sprintf(`select hashed_password from users where username='%s'`, username)
	row := DB.QueryRow(req)
	var hashedPassword string

	err := row.Scan(&hashedPassword)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}

	if HashPassword(username, password) != hashedPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\"%s\"", token)
}
