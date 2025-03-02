package swagger

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

func badRegisterRequest(w http.ResponseWriter, field string, reason string) {
	w.WriteHeader(http.StatusBadRequest)
	resp := InlineResponse400{field, reason}
	res, _ := json.Marshal(resp)
	fmt.Fprint(w, string(res))
}

const specialCharacters = "!@#$%^&*()'\""

func checkPasswordStrength(w http.ResponseWriter, password string) bool {
	if len(password) < 8 {
		badRegisterRequest(w, "Password must be at least 8 characters", password)
		return false
	}
	if len(password) > 32 {
		badRegisterRequest(w, "Password must be no more than 32 characters", password)
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
		badRegisterRequest(w, "Password must contain at least 1 digit", password)
		return false
	}
	if !has_upper {
		badRegisterRequest(w, "Password must contain at least 1 uppercase letter", password)
		return false
	}
	if !has_special {
		badRegisterRequest(w, "Password must contain at least 1 special character", password)
		return false
	}
	return true
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
	if !checkPasswordStrength(w, password) {
		return
	}
	email := r.URL.Query().Get("email")
	_, err := mail.ParseAddress(email)
	if err != nil {
		badRegisterRequest(w, "email", "Field email is invalid")
		return
	}

	log.Printf("Register:\nusername: %s\npassword: %s\nemail: %s\n", username, password, email)

	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	passwordHash.Write([]byte(username))

	req := fmt.Sprintf(`insert into users (username, email, hashed_password)
values ('%s', '%s', '%s')`, username, email, base64.URLEncoding.EncodeToString(passwordHash.Sum(nil)))
	_, err = DB.Query(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
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
