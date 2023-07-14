package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo!")
	}()

	func(text string) {
		fmt.Println(text)
	}("Parameter...") // This parameter is passed to anonymous function

	result := func(text string) string {
		return fmt.Sprintf("Received -> %s", text) // %s is string that will be replaced by variable
	}("Anonymous function with return")

	fmt.Println(result)
}
