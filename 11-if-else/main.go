package main

import "fmt"

func main() {
	number := 15

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less than 15")
	}

	// if init, the anotherNumbe variable is scoped to the statement
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Another number is greater than 0")
	} else if anotherNumber < -10 {
		fmt.Println("Another number is less than -10")
	} else {
		fmt.Println("Another number is between 0 and -10")
	}
}
