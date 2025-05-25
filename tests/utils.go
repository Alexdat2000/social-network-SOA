package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/lib/pq"
)

func isDocker() bool {
	val, exists := os.LookupEnv("DOCKER_RUNTIME")
	return exists && val == "1"
}

func getPostgresDb(host string, port int, dbname string) *sql.DB {
	if !isDocker() {
		host = "localhost"
	} else {
		port = 5432
	}
	dsn := fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=%s sslmode=disable", host, port, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}

func GetClickDb() *sql.DB {
	var dsn string
	if isDocker() {
		dsn = "clickhouse://default:clickhouse@clickhouse_stats:9000/default"
	} else {
		dsn = "clickhouse://default:clickhouse@localhost:9000/default"
	}
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	return db
}

func ClearTablePostgres(host string, port int, dbname, tablename string) {
	db := getPostgresDb(host, port, dbname)
	_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tablename))
	if err != nil {
		log.Fatalf("failed to truncate users table: %v", err)
	}
}

func CalcRowsInTable(host string, port int, dbname, tablename string) int {
	db := getPostgresDb(host, port, dbname)

	var count int
	row := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tablename))
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func ClearTableClick(tablename string) {
	db := GetClickDb()
	_, err := db.ExecContext(context.Background(), `ALTER TABLE `+tablename+` DELETE WHERE 1=1`)
	if err != nil {
		log.Fatal(err)
	}
}

func ExecuteInClick(reqs []string) {
	db := GetClickDb()
	for _, req := range reqs {
		_, err := db.ExecContext(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CalcRowsInClick(tablename string) int {
	db := GetClickDb()

	var count int
	row := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tablename))
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func SendRequestWithUrl(url, method, jwt, body string) (int, string) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Print(err.Error())
		return -1, ""
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}
	//command, _ := http2curl.GetCurlCommand(req)
	//fmt.Println(command)
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

func SendRequest(url, method, jwt, body string) (int, string) {
	if isDocker() {
		url = "http://gateway:8080" + url
	} else {
		url = "http://localhost:8080" + url
	}
	return SendRequestWithUrl(url, method, jwt, body)
}
