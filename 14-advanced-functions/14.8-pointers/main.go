package main

import "fmt"

// Functions always clone the variables that are passed as a parameter, not modifying them
func invertSignal(number int) int {
	return number * -1
}

// This function is receiving a memory reference, not the value
func invertSignalWithPointer(number *int) {
	fmt.Println(number)    // Memory address
	*number = *number * -1 // Value
}

func main() {
	number := 20
	result := invertSignal(number)
	fmt.Println(result)

	invertSignalWithPointer(&number)
	fmt.Println(number)
	fmt.Println(&number) // Memory address
}
