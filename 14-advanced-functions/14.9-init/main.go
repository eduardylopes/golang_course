package main

import "fmt"

var n int

// This function is executed before main, can be used only one per file.
func init() {
	fmt.Println("Executing function init")
	n = 10
}

func main() {
	fmt.Println("Executing function main")
	fmt.Println(n)
}
