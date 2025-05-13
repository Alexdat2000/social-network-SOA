package tests

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func SendRequest(url, method, jwt, body string) (int, string) {
	req, err := http.NewRequest(method, "http://localhost:8081"+url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		println(err.Error())
		return -1, ""
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return -1, ""
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return -1, ""
	}

	return resp.StatusCode, strings.TrimSpace(string(respBody))
}

func TestRegistration(t *testing.T) {
	status, resp := SendRequest("/users/login", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 401, status)
	assert.Equal(t, "User not found", resp)
}
