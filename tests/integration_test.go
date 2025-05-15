package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func ClearTable() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable"

	// Open database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Verify connection is alive
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Truncate the users table, reset identity, cascade to dependent tables
	_, err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	if err != nil {
		log.Fatalf("failed to truncate users table: %v", err)
	}
}

func SendRequest(url, method, jwt, body string) (int, string) {
	req, err := http.NewRequest(method, "http://localhost:8080"+url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Print(err.Error())
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
		log.Print(err.Error())
		return -1, ""
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		return -1, ""
	}

	return resp.StatusCode, strings.TrimSpace(string(respBody))
}

func TestRegistration(t *testing.T) {
	ClearTable()

	// Trying logging into non-existent user: error
	status, resp := SendRequest("/users/login", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 401, status)
	assert.Equal(t, "User not found", resp)

	// Registration without email: error
	status, resp = SendRequest("/users", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 400, status)
	assert.Equal(t, "Email is required", resp)

	// Successful registration
	status, resp = SendRequest("/users", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd",
  "email": "alex@example.com"
}`)
	assert.Equal(t, 200, status)
	jwt := resp[8 : len(resp)-2]
	assert.Equal(t, `{"jwt":"`+jwt+`"}`, resp)

	// Successful logging in
	status, resp = SendRequest("/users/login", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 200, status)
	newJwt := resp[8 : len(resp)-2]
	assert.Equal(t, `{"jwt":"`+newJwt+`"}`, resp)
	assert.Equal(t, jwt, newJwt)

	// Successful JWT check
	//status, resp = SendRequest("/users/auth", "GET", jwt, ``)
	//assert.Equal(t, 200, status)
	//assert.Equal(t, `{"username":"Alex"}`, resp)

	// Getting profile
	status, resp = SendRequest("/users/Alex", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	var result map[string]interface{}
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, "Alex", result["username"].(string))
	assert.Equal(t, "alex@example.com", result["email"].(string))

	// Updating profile
	status, resp = SendRequest("/users", "PATCH", jwt, `{
  "email": "alex2@example.com",
  "dateOfBirth": "1970-12-31",
  "phoneNumber": "800-555-0100"
}`)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, "Alex", result["username"].(string))
	assert.Equal(t, "alex2@example.com", result["email"].(string))
	assert.Equal(t, "1970-12-31", result["dateOfBirth"].(string))
	assert.Equal(t, "8005550100", result["phoneNumber"].(string))
}
