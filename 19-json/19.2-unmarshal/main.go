package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
}

func main() {
	dogJSON := `{"name":"Rex","race":"Pitbull","age":10}`
	var dog Dog

	if err := json.Unmarshal([]byte(dogJSON), &dog); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog)

	dogJSON2 := `{"name":"Bob","race":"Husky"}`
	var dog2 map[string]string

	if err := json.Unmarshal([]byte(dogJSON2), &dog2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog2)
}
