package main

import "fmt"

func checkApproval(n1, n2 float32) bool {
	// This will be called when the content of the scope are processed and before the return statement is executed
	defer fmt.Println("Average calculated. Result will be returned.")

	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	isApproved := checkApproval(7, 8)
	fmt.Println(isApproved)
	fmt.Println("Here...")
}
