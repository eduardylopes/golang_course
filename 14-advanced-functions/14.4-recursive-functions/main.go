package main

import (
	"fmt"
)

/*
	fibonacci = [1, 1, 2, 3, 5, 8, 13, 21...]

	fibonacci(1) = 1

	fibonacci(2) = 2

	fibonacci(3) = 3
		fibonacci(2) = 2
		fibonacci(1) = 1

	fibonacci(4) = 5
		fibonacci(3) = 3
		fibonacci(2) = 2

	fibonacci(5) = 8
		fibonacci(4) = 5
		fibonacci(3) = 3

	fibonacci(6) = 13
		fibonacci(5) = 8
		fibonacci(4) = 5
*/

func fibonacci(number uint) uint {
	if number < 2 {
		return number
	}

	return fibonacci(number-2) + fibonacci(number-1)
}

func main() {
	var sequence [28]uint

	for index := range sequence {
		sequence[index] = fibonacci(uint(index) + 1)
	}

	fmt.Println(sequence)
}
