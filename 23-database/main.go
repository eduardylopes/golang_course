package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionURL := "root:golang@/golang?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connectionURL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection is open")

	lines, err := db.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()

	fmt.Println(lines)
}
