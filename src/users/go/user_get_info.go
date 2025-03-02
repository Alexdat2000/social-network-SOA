package swagger

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func UsersGet(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("username")
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := fmt.Sprintf(`select id, username, email, first_name, last_name, date_of_birth, phone_number from users where username='%s'`, user)
	row := DB.QueryRow(req)

	nullableAnswer := struct {
		Id          int32
		Username    string
		Email       string
		FirstName   sql.NullString
		LastName    sql.NullString
		DateOfBirth sql.NullString
		PhoneNumber sql.NullString
	}{}

	err := row.Scan(&nullableAnswer.Id, &nullableAnswer.Username, &nullableAnswer.Email,
		&nullableAnswer.FirstName, &nullableAnswer.LastName, &nullableAnswer.DateOfBirth,
		&nullableAnswer.PhoneNumber)
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error when reading row: %v", err)
		return
	}

	ans := Profile{
		Id:       nullableAnswer.Id,
		Username: nullableAnswer.Username,
		Email:    nullableAnswer.Email,
	}

	if nullableAnswer.FirstName.Valid {
		ans.FirstName = nullableAnswer.FirstName.String
	}
	if nullableAnswer.LastName.Valid {
		ans.FirstName = nullableAnswer.LastName.String
	}
	if nullableAnswer.DateOfBirth.Valid {
		ans.FirstName = nullableAnswer.DateOfBirth.String
	}
	if nullableAnswer.PhoneNumber.Valid {
		ans.FirstName = nullableAnswer.PhoneNumber.String
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(ans)
	fmt.Fprint(w, string(res))
}
