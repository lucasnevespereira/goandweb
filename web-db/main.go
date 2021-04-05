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
CREATE TABLE IF NOT EXISTS books
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	author TEXT,
	page_count INT
)
`

var insertSchema = `
INSERT INTO "public"."books"("title", "author", "page_count") VALUES('Harry Potter', 'JK Rowling', 768);
`

func main() {
	pgConString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username, password, hostname, hostPort, databaseName)

	db, err := sql.Open("postgres", pgConString)
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

	insertStmt, err := db.Prepare(insertSchema)
	if err != nil {
		log.Fatal(err)
	}

	insertStmt.Exec()

	fmt.Println("Inserted document")

}
