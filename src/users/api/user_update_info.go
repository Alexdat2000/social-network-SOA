package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func SanitizePhone(phone string) (bool, string) {
	phone = strings.Replace(phone, " ", "", -1)
	phone = strings.Replace(phone, "(", "", -1)
	phone = strings.Replace(phone, ")", "", -1)
	phone = strings.Replace(phone, "-", "", -1)
	phone = strings.Replace(phone, "+", "", -1)

	ok := true
	for c := range phone {
		if !('0' <= phone[c] && phone[c] <= '9') {
			ok = false
			break
		}
	}
	return ok && len(phone) <= 13, phone
}

func (s Server) PatchUsers(w http.ResponseWriter, r *http.Request) {
	user, ok := ensureAuth(w, r)
	if !ok {
		return
	}

	var req PatchUsersJSONBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	if req.FirstName != nil {
		if len(*req.FirstName) < 1 || len(*req.FirstName) > 32 {
			http.Error(w, "First name must be between 1 and 32 characters.", http.StatusBadRequest)
			return
		}
	}
	if req.LastName != nil {
		if len(*req.LastName) < 1 || len(*req.LastName) > 32 {
			http.Error(w, "Last name must be between 1 and 32 characters.", http.StatusBadRequest)
			return
		}
	}
	if req.PhoneNumber != nil {
		ok, sanitized := SanitizePhone(*req.PhoneNumber)
		if !ok {
			http.Error(w, "Invalid phone number", http.StatusBadRequest)
			return
		}
		req.PhoneNumber = &sanitized
	}

	updates := User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		LastEditedAt: time.Now(),
	}
	if req.Email != nil {
		updates.Email = string(*req.Email)
	}
	if req.DateOfBirth != nil {
		updates.DateOfBirth = &req.DateOfBirth.Time
	}
	err = DB.Model(&User{}).
		Where("username = ?", user).
		Updates(updates).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error updating user: %v", err)
		return
	} else {
		s.GetUsersUsername(w, r, user)
	}
}
