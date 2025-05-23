package tests

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
