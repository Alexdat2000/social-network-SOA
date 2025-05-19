package api

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (s Server) PostUsersLogin(w http.ResponseWriter, r *http.Request) {
	var req PostUsersLoginJSONBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.Username) > 32 || len(req.Password) > 32 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var hashedPassword string
	err = s.DB.Model(&User{}).
		Select("hashed_password").
		Where("username = ?", req.Username).
		Take(&hashedPassword).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}

	if HashPassword(req.Username, req.Password) != hashedPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(s.Handlers, req.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error when creating token: %v", err)
		return
	} else {
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
	}
}
