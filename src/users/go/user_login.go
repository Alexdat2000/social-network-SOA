package swagger

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	req := fmt.Sprintf(`select id, hashed_password from users where username='%s'`, username)
	row := DB.QueryRow(req)
	var id int
	var hashedPassword string

	err := row.Scan(&id, &hashedPassword)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}

	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	passwordHash.Write([]byte(strconv.Itoa(id)))
	if base64.URLEncoding.EncodeToString(passwordHash.Sum(nil)) != hashedPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(strconv.Itoa(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\"%s\"", token)
}
