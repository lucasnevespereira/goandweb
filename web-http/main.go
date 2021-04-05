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

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "login.html")
		return

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Parsed Form method failed. err=%v \n", err)
			return
		}

		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "Go" && password == "test" {
			fmt.Fprintf(w, "You are logged in \n")
		} else {
			fmt.Fprintf(w, "Wring username / password \n")
		}
	}
}

func main() {

	// Assign a route to a func
	http.HandleFunc("/", hello)
	http.HandleFunc("/search", search)
	http.HandleFunc("/login", login)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
