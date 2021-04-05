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

func main() {

	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
