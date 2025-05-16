package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func Auth(w http.ResponseWriter, r *http.Request) (bool, string) {
	req, err := http.NewRequest("GET", "http://localhost:8080/users/auth", bytes.NewBuffer(nil))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, ""
	}
	req.Header = r.Header.Clone()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, ""
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false, ""
	}
	if resp.StatusCode != http.StatusOK {
		http.Error(w, strings.TrimSpace(string(respBody)), resp.StatusCode)
		return false, ""
	}
	var res map[string]string
	err = json.Unmarshal(respBody, &res)
	return true, res["username"]
}
