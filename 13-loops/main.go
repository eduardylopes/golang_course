package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementing i", i)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementing j", j)
		time.Sleep(time.Millisecond * 200)
	}

	names := [3]string{"Eduardy", "Ana", "Fabiano"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// If index is not needed use underscore notation "_"
	for _, name := range names {
		fmt.Println(name)
	}

	// Iterate a string
	for index, word := range "Eduardy" {
		fmt.Println(index, string(word)) // This iteration return the words in ASCII format, to get real word we need wrap the value using "string(value)"
	}

	user := map[string]string{
		"name":     "Eduardy",
		"lastName": "Morais",
	}

	// Iterate a map
	for key, value := range user {
		fmt.Println(key, value)
	}

	// Infinity Loop
	for {
		fmt.Println("Infinity loop")
		time.Sleep(time.Millisecond * 200)
	}
}
