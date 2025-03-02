package swagger

import (
	"encoding/json"
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

	req := fmt.Sprintf(`select * from users where username='%s'`, user)
	rows, err := DB.Query(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ans := Profile{}

	for rows.Next() {
		err := rows.Scan(&ans.Id, &ans.Username, &ans.Email, &ans.FirstName, &ans.LastName, &ans.DateOfBirth, &ans.PhoneNumber)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if ans.Username == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(ans)
	fmt.Fprint(w, string(res))
}
