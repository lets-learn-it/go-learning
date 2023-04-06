package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/todoapp?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
