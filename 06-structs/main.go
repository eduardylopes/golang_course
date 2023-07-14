package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	country string
	city    string
	street  string
}

func main() {
	var u user
	fmt.Println(u)

	u.name = "Eduardy"
	u.age = 28
	fmt.Println(u)

	a := address{"Brazil", "Santa Bárbara", "R. Antônio Miranda Filho"}

	u2 := user{"Eduardy", 21, a}
	fmt.Println(u2)

	u3 := user{name: "Eduardy"}
	fmt.Println(u3)
}
