package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	user := User{"Eduardy Lopes de Morais", "eduardylopes@gmail.com"}
	templates.ExecuteTemplate(w, "index.html", user)
}

func main() {
	templates = template.Must(template.ParseGlob("21-html/*.html"))

	http.HandleFunc("/", home)

	port := 3000
	fmt.Printf("Listening on http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
