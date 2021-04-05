package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Log request info
	fmt.Printf("%v %v \n", r.Method, r.URL)
	// Write a message as an HTTP response
	fmt.Fprintf(w, "Hello There!")
}


func search(w http.ResponseWriter, r *http.Request) {
	// For example request: http://localhost:8000/search?t=go&p=1
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")
	fmt.Fprintf(w, "Searching for term '%v' in page %v. \n", t, p)
}

func main() {

	// Assign a route to a func
	http.HandleFunc("/", hello)
	http.HandleFunc("/search", search)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
