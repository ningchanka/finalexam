package database

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

// FUNCTION init() call automaticly
func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func Conn() *sql.DB {
	return db
}
