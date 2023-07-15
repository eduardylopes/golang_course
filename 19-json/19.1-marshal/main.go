package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint8  `json:"age"`
}

func main() {
	dog := Dog{"Maylon", "Vira-lata", 9}
	fmt.Println(dog)

	dogJSON, err := json.Marshal(dog)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dogJSON))

	dog2 := map[string]string{
		"name": "Toby",
		"raca": "Pitbull",
		"age":  "10",
	}

	dogJSON2, err := json.Marshal(dog2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dogJSON2))
}
