package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
	"web-golang/helper"
)

func NewDB() *sql.DB {
	connStr := "postgres://postgres:dckyms@localhost/go-test-db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}