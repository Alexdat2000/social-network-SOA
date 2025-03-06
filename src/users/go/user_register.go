package api

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

func badRegisterRequest(w http.ResponseWriter, field string, reason string) {
	w.WriteHeader(http.StatusBadRequest)
	resp := InlineResponse400{field, reason}
	res, _ := json.Marshal(resp)
	fmt.Fprint(w, string(res))
}

const specialCharacters = "!@#$%^&*()'\""

func CheckPasswordStrength(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters"
	}
	if len(password) > 32 {
		return false, "Password must be less than 32 characters"
	}
	has_digit := false
	has_upper := false
	has_special := false
	for _, c := range password {
		if c >= '0' && c <= '9' {
			has_digit = true
		}
		if c >= 'A' && c <= 'Z' {
			has_upper = true
		}
		if strings.Contains(specialCharacters, string(c)) {
			has_special = true
		}
	}
	if !has_digit {
		return false, "Password must contain at least 1 digit"
	}
	if !has_upper {
		return false, "Password must contain at least 1 uppercase letter"
	}
	if !has_special {
		return false, "Password must contain at least 1 special character"
	}
	return true, ""
}

func HashPassword(username, password string) string {
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	passwordHash.Write([]byte{0})
	passwordHash.Write([]byte(username))
	return base64.URLEncoding.EncodeToString(passwordHash.Sum(nil))
}

func UsersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	username := r.URL.Query().Get("username")
	if username == "" {
		badRegisterRequest(w, "username", "Field username is required")
		return
	}
	if len(username) < 3 || len(username) > 32 {
		badRegisterRequest(w, "username", "Username must be between 3 and 32 characters")
	}
	password := r.URL.Query().Get("password")
	if password == "" {
		badRegisterRequest(w, "password", "Field password is required")
		return
	}
	passwordStrength, errMsg := CheckPasswordStrength(password)
	if !passwordStrength {
		badRegisterRequest(w, "password", errMsg)
		return
	}
	email := r.URL.Query().Get("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		badRegisterRequest(w, "email", "Field email is invalid")
		return
	}

	t := time.Now()
	_, err = DB.Exec(`insert into users (username, email, hashed_password, created_at, last_edited_at)
values ($1, $2, $3, $4, $5)`,
		username, email, HashPassword(username, password), t, t)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "\"This username already exists\"")
	} else {
		token, err := CreateToken(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "\"%s\"", token)
	}
}
