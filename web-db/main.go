package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)


const (
	hostname = "localhost"
	hostPort = 5432
	username = "postgres"
	password = ""
	databaseName = "test"
)

var schema = `
 CREATE TABLE IF NOT EXISTS book
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	author TEXT,
	page_count INT
)
`

func main() {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostPort, hostname, username, password, databaseName)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(schema)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()

	fmt.Println("Ping to db successful!")
}
