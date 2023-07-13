package main

import "fmt"

type person struct {
	name string
	lastName string
	age uint8
	height float32
}

type student struct {
	person
	curso string
	college string
}

func main() {
	p1 := person{"Eduardy", "Morais", 28, 1.84}
	fmt.Println(p1)

	s1 := student{p1, "IT", "MIT" }
	fmt.Println(s1.college)
}