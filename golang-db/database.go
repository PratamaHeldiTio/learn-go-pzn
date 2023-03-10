package golang_db

import (
	"database/sql"
	"log"
	"time"
)

func GetConnection() *sql.DB {
	connStr := "postgres://pratama:mecandoit@localhost:5432/go_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
