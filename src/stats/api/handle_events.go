package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"
)

type Event struct {
	Author string `json:"author"`
	Date   string `json:"date"`
	PostID int    `json:"post_id"`
}

func HandleEvent(db *sql.DB, eventType, eventJson string) {
	var event Event
	err := json.Unmarshal([]byte(eventJson), &event)
	if err != nil {
		log.Printf("Error unmarshalling event: %v", err)
		return
	}
	date, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		log.Printf("Error reading date: %v", err)
		return
	}
	var query string
	switch eventType {
	case "views":
		query = "INSERT INTO stats.views (post_id, author, date) VALUES (?, ?, ?)"
	case "likes":
		query = "INSERT INTO stats.likes (post_id, author, date) VALUES (?, ?, ?)"
	case "comments":
		query = "INSERT INTO stats.comments (post_id, author, date) VALUES (?, ?, ?)"
	}
	_, err = db.ExecContext(context.Background(), query, event.PostID, event.Author, date)
	if err != nil {
		log.Fatal(err)
	}
}
