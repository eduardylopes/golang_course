package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable1++

	fmt.Println(variable1, variable2)

	// Pointer is a memory reference
	var variable3 int
	var pointer *int

	variable3 = 100
	// Get memory reference
	pointer = &variable3

	// Read value from memory reference
	fmt.Println(variable3, pointer)

	// Only changes the value of memory address, not the pointer
	variable3 = 150
	fmt.Println(variable3, pointer)

	fmt.Println(variable3, *pointer)
}
