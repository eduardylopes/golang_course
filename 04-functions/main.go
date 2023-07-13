package main

import "fmt"

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

func mathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	minus := n1 - n2
	
	return sum, minus
}

func main() {
	result := sum(10, 20)
	fmt.Println(result)

	var f1 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f1Return := f1("Text from f1")
	fmt.Println(f1Return)

	resultSum, resultMinus := mathCalcs(10, 15)
	fmt.Println(resultSum, resultMinus)

	resultSum2, _ := mathCalcs(10, 15)
	fmt.Println(resultSum2)
}