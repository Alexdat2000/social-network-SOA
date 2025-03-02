package swagger

import (
	"fmt"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

func UsersPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jwt := r.URL.Query().Get("jwt")
	if jwt == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Parameter jwt is required")
		return
	}
	fieldName := r.URL.Query().Get("fieldName")
	if fieldName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Parameter fieldName is required")
		return
	}
	newValue := r.URL.Query().Get("newValue")
	if newValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Parameter newValue is required")
		return
	}

	username, err := ValidateToken(jwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid token")
		return
	}

	if fieldName == "firstName" || fieldName == "lastName" {
		if len(newValue) > 100 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Length must not exceed 100 characters")
			return
		}
	} else if fieldName == "email" {
		email := r.URL.Query().Get("email")
		_, err := mail.ParseAddress(email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid email address")
			return
		}
	} else if fieldName == "phone" {
		newValue = strings.Replace(newValue, " ", "", -1)
		newValue = strings.Replace(newValue, "(", "", -1)
		newValue = strings.Replace(newValue, ")", "", -1)
		newValue = strings.Replace(newValue, "-", "", -1)
		newValue = strings.Replace(newValue, "+", "", -1)

		ok := true
		for c := range newValue {
			if !('0' <= newValue[c] && newValue[c] <= '9') {
				ok = false
				break
			}
		}
		if len(newValue) > 13 || !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid phone number")
			return
		}
	} else if fieldName == "dateOfBirth" {
		t, err := time.Parse("1/2/2006", "5/11/2023")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Date of birth must be in format DD.MM.YYYY")
			return
		}
		newValue = t.Format("01/02/2006")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid field name")
		return
	}

	println(username)
}
