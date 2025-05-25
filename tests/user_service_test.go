package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistration(t *testing.T) {
	var status int
	var resp string
	ClearTablePostgres("postgres_users", 5432, "users", "users")

	// Trying logging into non-existent user: error
	status, resp = SendRequestWithUrl("http://users:8080/users/login", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 401, status)
	assert.Equal(t, "User not found", resp)

	// Registration without email: error
	status, resp = SendRequestWithUrl("http://users:8080/users", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 400, status)
	assert.Equal(t, "Email is required", resp)

	// Successful registration
	status, resp = SendRequestWithUrl("http://users:8080/users", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd",
  "email": "alex@example.com"
}`)
	assert.Equal(t, 200, status)
	jwt := resp[8 : len(resp)-2]
	assert.Equal(t, `{"jwt":"`+jwt+`"}`, resp)

	// Successful logging in
	status, resp = SendRequestWithUrl("http://users:8080/users/login", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd"
}`)
	assert.Equal(t, 200, status)
	newJwt := resp[8 : len(resp)-2]
	assert.Equal(t, `{"jwt":"`+newJwt+`"}`, resp)
	assert.Equal(t, jwt, newJwt)
}

func TestJWTValidation(t *testing.T) {
	var status int
	var resp string
	ClearTablePostgres("postgres_users", 5432, "users", "users")

	// Successful registration
	status, resp = SendRequestWithUrl("http://users:8080/users", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd",
  "email": "alex@example.com"
}`)
	jwt := resp[8 : len(resp)-2]

	// Successful JWT check
	status, resp = SendRequestWithUrl("http://users:8080/users/auth", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	assert.Equal(t, `{"username":"Alex"}`, resp)
}

func TestProfileUpdate(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	ClearTablePostgres("postgres_users", 5432, "users", "users")

	// Successful registration
	status, resp = SendRequestWithUrl("http://users:8080/users", "POST", "", `{
  "username": "Alex",
  "password": "P@ssW0rd",
  "email": "alex@example.com"
}`)
	jwt := resp[8 : len(resp)-2]

	// Getting profile
	status, resp = SendRequestWithUrl("http://users:8080/users/Alex", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, "Alex", result["username"].(string))
	assert.Equal(t, "alex@example.com", result["email"].(string))

	// Updating profile
	status, resp = SendRequestWithUrl("http://users:8080/users", "PATCH", jwt, `{
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
