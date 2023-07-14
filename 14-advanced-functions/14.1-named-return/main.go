package main

import "fmt"

// Short way to return multiple values inside a function
func mathCalc(n1, n2 int) (sum int, minus int) {
	sum = n1 + n2
	minus = n1 - n2
	return
}

func main() {
	sum, minus := mathCalc(3, 5)
	fmt.Println(sum, minus)
}
