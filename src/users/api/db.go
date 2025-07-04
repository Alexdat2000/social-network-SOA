package api

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	Username       string     `gorm:"primaryKey;size:100;unique;not null"`
	Email          string     `gorm:"size:100;not null"`
	HashedPassword string     `gorm:"size:64;not null"`
	FirstName      *string    `gorm:"size:100"`
	LastName       *string    `gorm:"size:100"`
	DateOfBirth    *time.Time `gorm:"type:date"`
	PhoneNumber    *string    `gorm:"size:15"`
	CreatedAt      time.Time  `gorm:"not null"`
	LastEditedAt   time.Time  `gorm:"not null"`
}

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

func InitDB() *gorm.DB {
	connStr := getDBConnectionString()

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	if db == nil {
		log.Fatal("No database connection")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to auto migrate:", err)
		return nil
	} else {
		log.Println("Successfully connected to the database!")
		return db
	}
}
