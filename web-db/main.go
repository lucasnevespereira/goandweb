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

type Book struct {
	Id int
	Title string
	Author string
	pageCount int
}

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

	fmt.Println("Ping to db successful!")

	// Create table from schema
	stmt, err := db.Prepare(schema)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()


	// Insert Document
	//var insertSchema = `INSERT INTO "public"."books"("title", "author", "page_count") VALUES('Harry Potter', 'JK Rowling', 768);`
	//insertStmt, err := db.Prepare(insertSchema)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//insertStmt.Exec()
	//fmt.Println("Inserted document")


	// Insert with Prepared Statement
	//myBook := Book{Title: "Deep Work", Author: "Cal Newport", pageCount: 308,}
	//insertStatement, _ := db.Prepare("INSERT INTO books (title, author, page_count) VALUES (?,?,?)")
	//_, err = insertStatement.Exec(myBook.Title, myBook.Author, myBook.pageCount)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// map to values
	//var id int
	//var title string
	//var author string
	//var pageCount int
	//rows, _ := db.Query("SELECT * from books")
	//for rows.Next() {
	//	rows.Scan(&id, &title,&author,&pageCount)
	//	fmt.Printf("id=%v, title=%v, author=%v, pageCount=%v \n", id,title,author,pageCount)
	//}

	// map to struct
	rows, _ := db.Query("SELECT * FROM books")
	b := Book{}
	for rows.Next() {
		rows.Scan(&b.Id,&b.Title,&b.Author,&b.pageCount)
		fmt.Printf("book=%v \n", b)
	}

}
