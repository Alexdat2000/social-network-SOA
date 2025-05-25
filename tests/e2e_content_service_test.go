package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContentServiceE2E(t *testing.T) {
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

	// Create post on Alex
	status, resp = SendRequest("/posts", "POST", jwt, `{
  "title": "Example post title",
  "content": "Example post content",
  "isPrivate": false,
  "tags": [
    "work",
    "Golang"
  ]
}`)
	assert.Equal(t, 201, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(1), result["post_id"].(float64))
	assert.Equal(t, "Example post title", result["title"].(string))
	assert.Equal(t, "Example post content", result["description"].(string))
	assert.Equal(t, "Alex", result["author"].(string))
	assert.Equal(t, "work", result["tags"].([]interface{})[0].(string))
	assert.Equal(t, "Golang", result["tags"].([]interface{})[1].(string))

	// Create private post on Alex
	status, resp = SendRequest("/posts", "POST", jwt, `{
  "title": "Second post",
  "content": "2",
  "isPrivate": true
}`)
	assert.Equal(t, 201, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(2), result["post_id"].(float64))
	assert.Equal(t, "Second post", result["title"].(string))
	assert.Equal(t, "2", result["description"].(string))
	assert.Equal(t, true, result["is_private"].(bool))

	// Create post on Alex2
	status, resp = SendRequest("/posts", "POST", jwt2, `{
  "title": "Three",
  "content": "3"
}`)
	assert.Equal(t, 201, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(3), result["post_id"].(float64))
	assert.Equal(t, "Three", result["title"].(string))
	assert.Equal(t, "3", result["description"].(string))
	assert.Equal(t, "Alex2", result["author"].(string))

	// Get 3 on Alex - should succeed
	status, resp = SendRequest("/posts/3", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(3), result["post_id"].(float64))
	assert.Equal(t, "Three", result["title"].(string))
	assert.Equal(t, "3", result["description"].(string))
	assert.Equal(t, "Alex2", result["author"].(string))

	// Get 2 on Alex2 - no access
	status, resp = SendRequest("/posts/2", "GET", jwt2, ``)
	assert.Equal(t, 403, status)
	assert.Equal(t, "no access to this private post", resp)

	// Update post
	status, resp = SendRequest("/posts/3", "PUT", jwt2, `{
  "content": "4",
  "tags": ["updated"]
}`)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(3), result["post_id"].(float64))
	assert.Equal(t, "Three", result["title"].(string))
	assert.Equal(t, "4", result["description"].(string))
	assert.Equal(t, "Alex2", result["author"].(string))
	assert.Equal(t, "updated", result["tags"].([]interface{})[0].(string))

	// Get list of posts
	status, resp = SendRequest("/posts?page=1", "GET", jwt, ``)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(2), result["total_pages"].(float64))
	assert.Equal(t, float64(1), result["post_ids"].([]interface{})[0].(float64))
	assert.Equal(t, float64(2), result["post_ids"].([]interface{})[1].(float64))

	// Get second page of posts
	status, resp = SendRequest("/posts?page=2", "GET", jwt, ``)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(2), result["total_pages"].(float64))
	assert.Equal(t, float64(3), result["post_ids"].([]interface{})[0].(float64))

	// Delete post
	status, resp = SendRequest("/posts/1", "DELETE", jwt, ``)
	assert.Equal(t, 204, status)

	// Get list of posts
	status, resp = SendRequest("/posts?page=1", "GET", jwt, ``)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(1), result["total_pages"].(float64))
	assert.Equal(t, float64(2), result["post_ids"].([]interface{})[0].(float64))
	assert.Equal(t, float64(3), result["post_ids"].([]interface{})[1].(float64))
}
