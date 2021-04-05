package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Address struct {
	Street string `json:"street"`
	City string `json:"city"`
	Country string `json:"country,omitempty"`

}

type User struct {
	Name string `json:"name"`
	Password string `json:"-"`
	Email string `json:"email"`
	Address Address `json:"address"`
}

func users(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			Name: "Bob",
			Password: "secret",
			Email: "bob@golang.org",
			Address: Address{
				Street: "15 rue Hade",
				City: "Paris",
				Country: "France",
			},
		},
		{
			Name: "Alice",
			Password: "test-secret",
			Email: "alice@golang.org",
			Address: Address{
				Street: "42 rue Elle",
				City: "Paris",
				Country: "",
			},
		},
	}

	// Encode Users in JSON
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

func main() {

	http.HandleFunc("/users", users)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
