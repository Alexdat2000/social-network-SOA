package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInteractionsE2E(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	ClearTablePostgres("postgres_users", 5432, "users", "users")
	ClearTablePostgres("postgres_content", 5433, "content", "entries")
	ClearTablePostgres("postgres_content", 5433, "content", "comments")

	// Register 2 users
	status, resp = SendRequest("/users", "POST", "", `{
 "username": "Alex",
 "password": "P@ssW0rd",
 "email": "alex@example.com"
}`)
	jwt := resp[8 : len(resp)-2]

	status, resp = SendRequest("/users", "POST", "", `{
 "username": "Alex2",
 "password": "P@ssW0rd",
 "email": "alex2@example.com"
}`)
	jwt2 := resp[8 : len(resp)-2]

	// Create posts
	_, _ = SendRequest("/posts", "POST", jwt, `{
 "title": "1",
 "content": "1"
}`)
	_, _ = SendRequest("/posts", "POST", jwt, `{
 "title": "2",
 "content": "2",
 "isPrivate": true
}`)
	_, _ = SendRequest("/posts", "POST", jwt2, `{
 "title": "3",
 "content": "3"
}`)

	// Like post
	status, resp = SendRequest("/posts/1/likes", "POST", jwt, ``)
	assert.Equal(t, 201, status)

	// Like own private post
	status, resp = SendRequest("/posts/2/likes", "POST", jwt, ``)
	assert.Equal(t, 201, status)

	// Like others private post - error
	status, resp = SendRequest("/posts/2/likes", "POST", jwt2, ``)
	assert.Equal(t, 403, status)

	// Comment on post
	status, resp = SendRequest("/posts/1/comments", "POST", jwt, `{"text": "Comment 1"}`)
	assert.Equal(t, 201, status)
	status, resp = SendRequest("/posts/1/comments", "POST", jwt2, `{"text": "Comment 2"}`)
	assert.Equal(t, 201, status)
	status, resp = SendRequest("/posts/1/comments", "POST", jwt, `{"text": "Comment 3"}`)
	assert.Equal(t, 201, status)

	// Comment on others private post - error
	status, resp = SendRequest("/posts/2/comments", "POST", jwt2, `{"text": "Comment 4"}`)
	assert.Equal(t, 403, status)

	// Get comments
	status, resp = SendRequest("/posts/1/comments?page=1", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(2), result["total_pages"].(float64))
	assert.Equal(t, float64(1), result["comments"].([]interface{})[0].(map[string]interface{})["id"].(float64))
	assert.Equal(t, "Alex", result["comments"].([]interface{})[0].(map[string]interface{})["author"].(string))
	assert.Equal(t, "Comment 1", result["comments"].([]interface{})[0].(map[string]interface{})["text"].(string))
	assert.Equal(t, float64(2), result["comments"].([]interface{})[1].(map[string]interface{})["id"].(float64))
	assert.Equal(t, "Alex2", result["comments"].([]interface{})[1].(map[string]interface{})["author"].(string))
	assert.Equal(t, "Comment 2", result["comments"].([]interface{})[1].(map[string]interface{})["text"].(string))
}
