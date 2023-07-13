package main

import "fmt"

func main() {
	// Arithmetics
	sum := 1 + 2
	minus := 3 - 2
	times := 4 * 2
	dividedBy := 5 / 2
	remainder := 5 % 2
	fmt.Println(sum, minus, dividedBy, times, remainder)

	// Assignment
	var variable1 string = "variable1"
	variable2 := "variable2"
	fmt.Println(variable1, variable2)

	// Relational Operators
	fmt.Println(1 > 2)
	fmt.Println(1 >- 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Logical Operators
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(!true)

	// Unary Operators
	number := 10
	number++
	number += 15
	number--
	number -= 1
	number *= 3
	fmt.Println(number)

	var text string

	// Isn't possible use ternary operators in golang
	if number > 5 {
		text = "Greater than 5"
	} else {
		text = "Less than 5"
	}

	fmt.Println(text)
}