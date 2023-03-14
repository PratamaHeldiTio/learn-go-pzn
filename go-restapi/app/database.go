package app

import (
	"database/sql"
	"go-restapi/helper"
	"time"
)

func NewDB() *sql.DB {
	connStr := "postgres://pratamadev:tamadev99@localhost:5432/go_restapi?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
