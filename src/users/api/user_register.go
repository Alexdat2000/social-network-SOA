package api

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const specialCharacters = "!@#$%^&*()"

func CheckPasswordStrength(password string) string {
	if len(password) < 8 {
		return "Password must be at least 8 characters"
	}
	if len(password) > 32 {
		return "Password must be less than 32 characters"
	}
	hasDigit := false
	hasUpper := false
	hasSpecial := false
	for _, c := range password {
		if c >= '0' && c <= '9' {
			hasDigit = true
		}
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
		if strings.Contains(specialCharacters, string(c)) {
			hasSpecial = true
		}
		if !('0' <= c && c <= '9') && !('a' <= c && c <= 'z') && !('A' <= c && c <= 'Z') && !strings.Contains(specialCharacters, string(c)) {
			return "Illegal character " + string(c)
		}
	}
	if !hasDigit {
		return "Password must contain at least 1 digit"
	}
	if !hasUpper {
		return "Password must contain at least 1 uppercase letter"
	}
	if !hasSpecial {
		return "Password must contain at least 1 special character"
	}
	return ""
}

func HashPassword(username, password string) string {
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	passwordHash.Write([]byte{0})
	passwordHash.Write([]byte(username))
	return base64.URLEncoding.EncodeToString(passwordHash.Sum(nil))
}

func (s Server) PostUsers(w http.ResponseWriter, r *http.Request) {
	var req PostUsersJSONBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	if len(req.Username) < 3 || len(req.Password) > 32 {
		http.Error(w, "Username must be between 3 and 32 characters", http.StatusBadRequest)
		return
	}
	passwordBadReason := CheckPasswordStrength(req.Password)
	if passwordBadReason != "" {
		http.Error(w, passwordBadReason, http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	hashedPassword := HashPassword(req.Username, req.Password)
	t := time.Now()
	user := User{
		Username:       req.Username,
		Email:          string(req.Email),
		HashedPassword: hashedPassword,
		CreatedAt:      t,
		LastEditedAt:   t,
	}
	err = DB.Create(&user).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		http.Error(w, "User already exists", http.StatusBadRequest)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating user: %v", err)
	} else {
		token, err := CreateToken(req.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := map[string]string{
			"jwt": token,
		}
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error while writing json: %v", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonBytes)
		err = ReportRegisterToKafka(req.Username, string(req.Email), t)
		if err != nil {
			log.Printf("Error while reporting register to kafka: %v", err)
		}
	}
}
