package api

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func getDBConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSL_MODE")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}

var DB *sql.DB

func connectToDB() {
	connStr := getDBConnectionString()

	maxRetries := 5
	var err error
	for i := 0; i < maxRetries; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Failed to open DB connection: %v", err)
			time.Sleep(time.Second * 3)
			continue
		}

		err = DB.Ping()
		if err == nil {
			break
		}

		log.Printf("Failed to ping DB (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 3)
	}
	if err != nil {
		log.Fatalf("Failed to connect to the database after %d attempts: %v", maxRetries, err)
	}
}

func createTables() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS entries
(
    id              INT UNIQUE NOT NULL PRIMARY KEY,
    title           TEXT NOT NULL,
    description     TEXT NOT NULL,
    author          VARCHAR(64) NOT NULL,
    created_at      TIMESTAMP,
    last_edited_at  TIMESTAMP,
    is_private      boolean,
    tags            TEXT
);
`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func InitDB() {
	connectToDB()
	log.Println("Successfully connected to the database!")
	createTables()
	log.Println("Successfully set up tables!")
}
