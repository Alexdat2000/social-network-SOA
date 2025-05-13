package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

var ErrNoToken = errors.New("no token provided")
var ErrInvalidToken = errors.New("invalid token")
var ErrUserNotFound = errors.New("user not found")

func extractTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoToken
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", ErrInvalidToken
	}

	token := strings.TrimSpace(authHeader[len(prefix):])
	if token == "" {
		return "", ErrInvalidToken
	}

	return token, nil
}

func auth(r *http.Request) (string, error) {
	jwt, err := extractTokenFromRequest(r)
	name, err := ValidateToken(jwt)
	if err != nil {
		return "", err
	}

	var totalCount int64
	err = DB.Model(&User{}).
		Where("username = ?", name).
		Count(&totalCount).Error
	if err != nil {
		return "", err
	}
	if totalCount == 0 {
		return "", ErrUserNotFound
	}
	return name, nil
}

func ensureAuth(w http.ResponseWriter, r *http.Request) (string, bool) {
	name, err := auth(r)
	if errors.Is(err, ErrNoToken) || errors.Is(err, ErrInvalidToken) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return "", false
	} else if errors.Is(err, ErrUserNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return "", false
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", false
	} else {
		return name, true
	}
}

func (s Server) GetUsersAuth(w http.ResponseWriter, r *http.Request) {
	name, err := auth(r)
	if errors.Is(err, ErrNoToken) || errors.Is(err, ErrInvalidToken) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} else if errors.Is(err, ErrUserNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		data := map[string]string{
			"username": name,
		}
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error while writing json: %v", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonBytes)
	}
}
