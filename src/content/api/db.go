package api

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func getDBConnectionString(dbname string) string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSL_MODE")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}

func InitDB(dbname string, scheme interface{}) *gorm.DB {
	connStr := getDBConnectionString(dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	if db == nil {
		log.Fatal("No database connection")
	}

	err = db.AutoMigrate(scheme)
	if err != nil {
		log.Fatal("failed to auto migrate:", err)
	}
	log.Printf("Successfully connected to the database " + dbname)
	return db
}
