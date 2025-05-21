package tests

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/lib/pq"
)

func ClearTablePostgres(port int, dbname, tablename string) {
	dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=postgres dbname=%s sslmode=disable", port, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_, err = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tablename))
	if err != nil {
		log.Fatalf("failed to truncate users table: %v", err)
	}
}

func ClearTableClick(tablename string) {
	dsn := "clickhouse://default:clickhouse@localhost:9000/default"
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	_, err = db.ExecContext(context.Background(), `ALTER TABLE `+tablename+` DELETE WHERE 1=1`)
	if err != nil {
		log.Fatal(err)
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

func TestUserService(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	ClearTablePostgres(5432, "users", "users")
	ClearTablePostgres(5433, "content", "entries")
	ClearTablePostgres(5433, "content", "comments")

	// Trying logging into non-existent user: error
	status, resp = SendRequest("/users/login", "POST", "", `{
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

func TestContentService(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	ClearTablePostgres(5432, "users", "users")
	ClearTablePostgres(5433, "content", "entries")
	ClearTablePostgres(5433, "content", "comments")

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

func TestInteractions(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	ClearTablePostgres(5432, "users", "users")
	ClearTablePostgres(5433, "content", "entries")
	ClearTablePostgres(5433, "content", "comments")

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

func TestStatsService(t *testing.T) {
	var status int
	var resp string
	var result map[string]interface{}
	var resultList []map[string]interface{}
	ClearTablePostgres(5432, "users", "users")
	ClearTablePostgres(5433, "content", "entries")
	ClearTablePostgres(5433, "content", "comments")
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
