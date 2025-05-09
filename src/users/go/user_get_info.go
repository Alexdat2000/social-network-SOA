package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func UsersGet(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("username")
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := fmt.Sprintf(`select email, first_name, last_name, date_of_birth, phone_number, created_at, last_edited_at
from users where username='%s'`, user)
	row := DB.QueryRow(req)

	nullableAnswer := struct {
		Email        string
		FirstName    sql.NullString
		LastName     sql.NullString
		DateOfBirth  sql.NullString
		PhoneNumber  sql.NullString
		CreatedAt    time.Time
		LastEditedAt time.Time
	}{}

	err := row.Scan(&nullableAnswer.Email,
		&nullableAnswer.FirstName, &nullableAnswer.LastName, &nullableAnswer.DateOfBirth,
		&nullableAnswer.PhoneNumber, &nullableAnswer.CreatedAt, &nullableAnswer.LastEditedAt)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}

	ans := Profile{
		Username:     user,
		Email:        nullableAnswer.Email,
		CreatedAt:    int32(nullableAnswer.CreatedAt.Unix()),
		LastEditedAt: int32(nullableAnswer.LastEditedAt.Unix()),
	}
	if nullableAnswer.FirstName.Valid {
		ans.FirstName = nullableAnswer.FirstName.String
	}
	if nullableAnswer.LastName.Valid {
		ans.LastName = nullableAnswer.LastName.String
	}
	if nullableAnswer.DateOfBirth.Valid {
		ans.DateOfBirth = nullableAnswer.DateOfBirth.String
	}
	if nullableAnswer.PhoneNumber.Valid {
		ans.PhoneNumber = nullableAnswer.PhoneNumber.String
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(ans)
	fmt.Fprint(w, string(res))
}
