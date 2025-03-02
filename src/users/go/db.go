package swagger

import (
	"database/sql"
	"fmt"
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

func Init() {
	connStr := getDBConnectionString()
	fmt.Printf("Connecting to PostgreSQL: %s\n", connStr)

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Failed to open DB connection: %v", err)
			time.Sleep(time.Second * 3)
			continue
		}

		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("Failed to ping DB (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 3)
	}
	log.Println("Successfully connected to the database!")
}
