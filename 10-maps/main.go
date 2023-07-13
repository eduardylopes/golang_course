package main

import "fmt"

func main() {
	user := map[string]string {
		"name": "Eduardy",
		"lastName": "Morais",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"firstName": "Eduardy",
			"lastName": "Morais",
		},
		"course": {
			"name": "IT",
			"college": "MIT",
		},
	}

	fmt.Println(user2)

	// Delete an element by key
	delete(user2, "name")
	fmt.Println(user2)

	// Insert a new element
	user2["name"] = map[string]string {
		"firstName": "Eduardy",
		"lastName": "Morais",
	}
	fmt.Println(user2)
}