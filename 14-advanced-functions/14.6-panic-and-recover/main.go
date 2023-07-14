package main

import "fmt"

func recoveryExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution recovered")
	}
}

func checkApproval(n1, n2 float32) bool {
	defer recoveryExecution()

	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("The average is 6")
}

func main() {
	isApproved := checkApproval(6, 6)
	fmt.Println(isApproved)
}
