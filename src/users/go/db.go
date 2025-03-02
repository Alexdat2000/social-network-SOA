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
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users
(
    username        VARCHAR(100) UNIQUE NOT NULL PRIMARY KEY,
    email           VARCHAR(100) NOT NULL,
    hashed_password VARCHAR(64)  NOT NULL,
    first_name      VARCHAR(100),
    last_name       VARCHAR(100),
    date_of_birth   DATE,
    phone_number    varchar(15)
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
