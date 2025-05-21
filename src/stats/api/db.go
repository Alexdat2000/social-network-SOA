package api

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func InitClick() *sql.DB {
	dsn := "clickhouse://default:clickhouse@clickhouse_stats:9000/default"
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to Click")
	return db
}
