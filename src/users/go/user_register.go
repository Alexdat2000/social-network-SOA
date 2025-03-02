package swagger

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/mail"
	"strconv"
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

	id := int(uuid.New().ID() / 2)
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	req := fmt.Sprintf(`insert into users (id, username, email, hashed_password)
values ('%d', '%s', '%s', '%s')`, id, username, email, base64.URLEncoding.EncodeToString(passwordHash.Sum(nil)))
	println(req)
	_, err = DB.Query(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
	} else {
		token, err := CreateToken(strconv.Itoa(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "\"%s\"", token)
	}
}
