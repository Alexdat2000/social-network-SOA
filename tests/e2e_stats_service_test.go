package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatsServiceE2E(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	var resultList []map[string]interface{}
	ClearTablePostgres("postgres_users", 5432, "users", "users")
	ClearTablePostgres("postgres_content", 5433, "content", "entries")
	ClearTablePostgres("postgres_content", 5433, "content", "comments")
	ClearTableClick("stats.views")
	ClearTableClick("stats.likes")
	ClearTableClick("stats.comments")

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
	_, _ = SendRequest("/posts", "POST", jwt2, `{
 "title": "2",
 "content": "2"
}`)

	// View post
	status, resp = SendRequest("/posts/1", "GET", jwt, ``)

	// Like posts
	status, resp = SendRequest("/posts/1/likes", "POST", jwt, ``)
	status, resp = SendRequest("/posts/2/likes", "POST", jwt, ``)
	status, resp = SendRequest("/posts/2/likes", "POST", jwt2, ``)

	// Comments on post
	status, resp = SendRequest("/posts/1/comments", "POST", jwt, `{"text": "Comment 1"}`)
	status, resp = SendRequest("/posts/1/comments", "POST", jwt2, `{"text": "Comment 2"}`)
	status, resp = SendRequest("/posts/1/comments", "POST", jwt, `{"text": "Comment 3"}`)

	// Get post stats
	status, resp = SendRequest("/posts/1/stats", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(2), result["views"].(float64))
	assert.Equal(t, float64(1), result["likes"].(float64))
	assert.Equal(t, float64(3), result["comments"].(float64))

	status, resp = SendRequest("/posts/2/stats", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &result)
	assert.Equal(t, float64(1), result["views"].(float64))
	assert.Equal(t, float64(2), result["likes"].(float64))

	// Get daily stats
	status, resp = SendRequest("/posts/1/stats/daily?metric=comments", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &resultList)
	assert.Equal(t, float64(3), resultList[0]["count"].(float64))

	// Get top posts
	status, resp = SendRequest("/posts/top10?metric=views", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &resultList)
	assert.Equal(t, float64(1), resultList[0]["post_id"].(float64))
	assert.Equal(t, float64(2), resultList[0]["count"].(float64))
	assert.Equal(t, float64(2), resultList[1]["post_id"].(float64))
	assert.Equal(t, float64(1), resultList[1]["count"].(float64))

	// Get top users
	status, resp = SendRequest("/users/top10?metric=likes", "GET", jwt, ``)
	assert.Equal(t, 200, status)
	_ = json.Unmarshal([]byte(resp), &resultList)
	assert.Equal(t, "Alex2", resultList[0]["username"].(string))
	assert.Equal(t, float64(2), resultList[0]["count"].(float64))
	assert.Equal(t, "Alex", resultList[1]["username"].(string))
	assert.Equal(t, float64(1), resultList[1]["count"].(float64))
}
