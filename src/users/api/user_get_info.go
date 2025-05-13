package api

import (
	"encoding/json"
	"errors"
	"fmt"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (s Server) GetUsersUsername(w http.ResponseWriter, r *http.Request, username string) {
	_, ok := ensureAuth(w, r)
	if !ok {
		return
	}

	var info User
	err := DB.Model(&User{}).
		Select("email, first_name, last_name, date_of_birth, phone_number, created_at, last_edited_at").
		Where("username = ?", username).
		Take(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}

	createdAt := int(info.CreatedAt.Unix())
	lastEditedAt := int(info.LastEditedAt.Unix())
	dateOfBirth := openapi_types.Date{Time: *info.DateOfBirth}
	ans := Profile{
		Username:     info.Username,
		Email:        openapi_types.Email(info.Email),
		CreatedAt:    &createdAt,
		LastEditedAt: &lastEditedAt,
		FirstName:    info.FirstName,
		LastName:     info.LastName,
		DateOfBirth:  &dateOfBirth,
		PhoneNumber:  info.PhoneNumber,
	}

	res, err := json.Marshal(ans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error when marshalling json: %v", err)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, string(res))
	}
}
