package main

import "fmt"

func closure() func() {
	text := "Inside from closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Inside the main function"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}
