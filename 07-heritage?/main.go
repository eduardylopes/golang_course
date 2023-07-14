package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      uint8
	height   float32
}

type Student struct {
	Person
	curso   string
	college string
}

func main() {
	person := Person{"Eduardy", "Morais", 28, 1.84}
	fmt.Println(person)

	student := Student{person, "IT", "MIT"}
	fmt.Println(student.college)
}
