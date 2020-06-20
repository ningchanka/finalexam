package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ningchanka/finalexam/customer"
)

func checkDatabaseConnectionAndTable(database_url string) {
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal("Fail to connect to database. ", err)
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS customer (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
		);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("create table success")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2019"
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		log.Fatal("Fail to get DATABASE_URL from environment. ")
	}

	checkDatabaseConnectionAndTable(database_url)

	r := customer.SetupRouter()
	r.Run(fmt.Sprintf(":%s", port))
}
