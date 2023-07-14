package main

import "fmt"

type User struct {
	name    string
	age     uint8
	address Address
}

type Address struct {
	country string
	city    string
	street  string
}

func main() {
	var user User
	fmt.Println(user)

	user.name = "Eduardy"
	user.age = 28
	fmt.Println(user)

	address := Address{"Brazil", "Santa Bárbara", "R. Antônio Miranda Filho"}

	user2 := User{"Eduardy", 21, address}
	fmt.Println(user2)

	user3 := User{name: "Eduardy"}
	fmt.Println(user3)
}
