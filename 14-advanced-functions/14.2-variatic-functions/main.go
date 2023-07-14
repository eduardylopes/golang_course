package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Is permitted only one variatic parameter per function, and it must be the last parameter
func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	result := sumAll(1, 2, 3, 4, 5)
	fmt.Println(result)

	write("Olá mundo", 1, 2, 3, 4, 5, 6, 7)
}
