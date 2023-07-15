package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users..."))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
