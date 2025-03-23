package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
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

func UsersPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jwt := r.URL.Query().Get("jwt")
	if jwt == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "\"Parameter jwt is required\"")
		return
	}
	fieldName := r.URL.Query().Get("fieldName")
	if fieldName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "\"Parameter fieldName is required\"")
		return
	}
	newValue := r.URL.Query().Get("newValue")
	if newValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "\"Parameter newValue is required\"")
		return
	}

	username, err := ValidateToken(jwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "\"Invalid token\"")
		return
	}

	var res sql.Result

	if fieldName == "firstName" || fieldName == "lastName" {
		if len(newValue) > 100 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "\"Length must not exceed 100 characters\"")
			return
		}
		if fieldName == "firstName" {
			res, err = DB.Exec(`update users set first_name = $1, last_edited_at=$2 where username = $3`, newValue, time.Now(), username)
		} else {
			res, err = DB.Exec(`update users set last_name = $1, last_edited_at=$2 where username = $3`, newValue, time.Now(), username)
		}
	} else if fieldName == "email" {
		email := r.URL.Query().Get("email")
		_, err := mail.ParseAddress(email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "\"Invalid email address\"")
			return
		}
		res, err = DB.Exec(`update users set email = $1, last_edited_at=$2 where username = $3`, newValue, time.Now(), username)
	} else if fieldName == "phone" {
		ok, newValue := SanitizePhone(newValue)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "\"Invalid phone number\"")
			return
		}
		res, err = DB.Exec(`update users set phone_number = $1, last_edited_at=$2 where username = $3`, newValue, time.Now(), username)
	} else if fieldName == "dateOfBirth" {
		t, err := time.Parse("2006-1-2", newValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "\"Date of birth must be in format YYYY-MM-DD\"")
			return
		}
		newValue = t.Format("2006-1-2")
		res, err = DB.Exec(`update users set date_of_birth = $1, last_edited_at=$2 where username = $3`, newValue, time.Now(), username)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "\"Invalid field name\"")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
		return
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
